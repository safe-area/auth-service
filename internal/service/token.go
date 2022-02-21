package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (s *service) createToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&Claims{
			jwt.StandardClaims{},
			userId,
		})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *service) parseToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.JWTSecret), nil
		})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	return claims.UserId, nil
}
