package helper

import (
	"TrainTracking/internal/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var (
	SecretKey = config.GetApp().JWTSecret
)

func CreateToken(email, role string) (string, error) {
	env := config.GetApp()
	jwtExpirationTime, _ := strconv.Atoi(env.JWTExpirationTime)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                                                                 // Subject (user identifier)
		"iss": config.GetApp().Name,                                                  // Issuer
		"aud": role,                                                                  // Audience (user role)
		"exp": time.Now().Add(time.Minute * time.Duration(jwtExpirationTime)).Unix(), // Expiration time
		"iat": time.Now().Unix(),                                                     // Issued at
	})

	tokenString, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
