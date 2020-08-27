package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func GenerateJWT(key, secret string, jwtLifetime int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtLifetime) * time.Minute)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    key, // key can be understood as username
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)

	return token, err
}

func ParseJWT(token, secret string) (*Claims, error) {
	jwtSecret := []byte(secret)

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
