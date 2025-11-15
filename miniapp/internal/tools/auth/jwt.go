package auth

import (
	"errors"
	"iskra/shared/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

func GetDataFromJWTToken(tokenString string, cfg *config.Config) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return cfg.JWT.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, ErrInvalidToken
	}

	// извлечение claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return 0, ErrTokenExpired
			}
		}

		if userID, ok := claims["userID"].(int); ok {
			return int64(userID), nil
		}
	}
	return 0, ErrInvalidToken
}

func GenJWTToken(userID int64, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Duration(cfg.JWT.TokenTTLSeconds)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.JWT.SecretKey)
}
