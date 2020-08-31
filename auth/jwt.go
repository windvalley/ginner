package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims payload infos
type Claims struct {
	jwt.StandardClaims
}

// GenerateJWT Parameter key can be understood as username or appkey.
func GenerateJWT(key, secret string, jwtLifetime int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtLifetime) * time.Second)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    key,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err
}

// ParseJWT get *Claims from jwt token
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

// GetPayload get payload from the second segment of jwt token
func GetPayload(seg string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}

	claimsBytes, err := jwt.DecodeSegment(seg)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(claimsBytes, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// GetSecretOfAppkey get appSecret by appKey
func GetSecretOfAppkey(appKey string) (string, bool) {
	userInfo, ok := userInfos[appKey]
	if !ok {
		return "", false
	}

	return userInfo.keySecret.JWT, true
}
