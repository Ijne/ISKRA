package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"iskra/shared/config"
	"iskra/shared/models"
	"iskra/shared/storage/memgraph"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SaveUserAvatar(userID int64, base64Image string) (string, error) {
	// Проверяем, является ли строка base64 изображением
	if !strings.HasPrefix(base64Image, "data:image") {
		return base64Image, nil
	}

	// Извлекаем данные
	parts := strings.SplitN(base64Image, ",", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid base64 image format")
	}

	// Обрабатываем MIME type
	mimePart := strings.TrimSuffix(strings.Split(parts[0], ";")[0], "")
	mimeType := strings.TrimPrefix(mimePart, "data:")

	var extension string
	switch {
	case strings.Contains(mimeType, "jpeg"), strings.Contains(mimeType, "jpg"):
		extension = "jpg"
	case strings.Contains(mimeType, "png"):
		extension = "png"
	case strings.Contains(mimeType, "gif"):
		extension = "gif"
	case strings.Contains(mimeType, "webp"):
		extension = "webp"
	default:
		return "", fmt.Errorf("unsupported image type: %s", mimeType)
	}

	// Декодируем base64
	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %v", err)
	}

	// Создаем директорию
	projectRoot, _ := os.Getwd()
	avatarsDir := filepath.Join(projectRoot, "static", "img")

	// Сохраняем файл
	filename := fmt.Sprintf("user_%d.%s", userID, extension)
	filepath := filepath.Join(avatarsDir, filename)

	if err := os.WriteFile(filepath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return fmt.Sprintf("/static/img/%s", filename), nil
}

func CreateUserHandler(cfg *config.Config, s *postgres.Storage, g *memgraph.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var user models.UserCreate
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] json err: %s\n", err)
				return
			}

			path, err := SaveUserAvatar(user.ID, user.Photo)
			if err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] SaveUserAvatart err: %s\n", err)
			}
			user.Photo = path

			if err := s.UserRepo.CreateUser(user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] psgrs err: %s\n", err)
				return
			}

			if err := g.SocialWebRepo.CreateUser(user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] memgrph err: %s\n", err)
				return
			}

		default:
			log.Printf("ERROR FROM[CrateUserHandler] Not allowed http method: %s", r.Method)
		}
	}
}
