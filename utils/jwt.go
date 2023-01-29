package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

var jwtSecret []byte

// GenerateToken 生成token
func GenerateToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        strconv.FormatUint(uint64(id), 10),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

// VerifyToken 校验jwt，返回userId
func VerifyToken(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Check if the token is valid
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}
	if err := claims.Valid(); err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(claims.ID, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
