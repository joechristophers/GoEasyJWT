package goeasyjwt

import (
	"errors"
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("eycvkjdhcusdavuyusydnmcs")

func GenerateToken(Claims map[string]any, jwtKey []byte, expireHour int) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(expireHour)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	maps.Copy(claims, Claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func VerifyToken(tokenString string, jwtKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("expired token")
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
