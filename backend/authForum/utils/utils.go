package utils

import (
	"authForum/envs"
	"authForum/logging"
	"authForum/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// checks the password hash using bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// token generation
func GenerateTokens(userID uint) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // Token validity period 12 hours
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(), // Valid for one week
	})
	signedAccessToken, err := accessToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))
	if err != nil {
		logging.LogError.Println("failed to sign AccessToken")
		return models.Tokens{}, err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))
	if err != nil {
		logging.LogError.Println("failed to sign RefreshToken")
		return models.Tokens{}, err
	}

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

// refreshToken validation
func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDValue, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("user_id claim is not a float64")
		}
		return uint(userIDValue), nil 
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}
