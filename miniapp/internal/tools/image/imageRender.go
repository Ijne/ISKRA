package image

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ConvertImageToBase64(imagePath string) (string, error) {
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return "", fmt.Errorf("image file not found: %s", imagePath)
	}

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %v", err)
	}

	var mimeType string
	ext := strings.ToLower(filepath.Ext(imagePath))
	switch ext {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".png":
		mimeType = "image/png"
	case ".gif":
		mimeType = "image/gif"
	case ".webp":
		mimeType = "image/webp"
	default:
		mimeType = http.DetectContentType(imageData)
		if !strings.HasPrefix(mimeType, "image/") {
			return "", fmt.Errorf("unsupported image format: %s", ext)
		}
	}
	base64String := base64.StdEncoding.EncodeToString(imageData)

	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64String)

	return dataURL, nil
}

func SaveUserAvatar(userID int64, base64Image string) (string, error) {
	if !strings.HasPrefix(base64Image, "data:image") {
		return base64Image, nil
	}

	parts := strings.SplitN(base64Image, ",", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid base64 image format")
	}

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

	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %v", err)
	}

	projectRoot, _ := os.Getwd()
	avatarsDir := filepath.Join(projectRoot, "static", "img")

	filename := fmt.Sprintf("user_%d.%s", userID, extension)
	filepath := filepath.Join(avatarsDir, filename)

	if err := os.WriteFile(filepath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return fmt.Sprintf("/static/img/%s", filename), nil
}

func GetUserPhoto(photoPath string) string {
	if photoPath == "" || photoPath == "null" || photoPath == "undefined" {
		return ""
	}

	if strings.HasPrefix(photoPath, "data:image") {
		return photoPath
	}

	cleanPath := strings.TrimPrefix(photoPath, "/")
	base64Photo, err := ConvertImageToBase64(cleanPath)
	if err != nil {
		fmt.Printf("Warning: could not convert photo to base64: %v\n", err)
		return ""
	}

	return base64Photo
}
