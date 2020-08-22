package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"use-gin/config"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	jwtSecret := []byte(config.Conf().Auth.JWTSecret)

	nowTime := time.Now()
	expireTime := nowTime.Add(
		time.Duration(config.Conf().Auth.JWTLifetime) * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Conf().AppName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseJWT(token string) (*Claims, error) {
	jwtSecret := []byte(config.Conf().Auth.JWTSecret)

	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
