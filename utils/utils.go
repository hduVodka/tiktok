package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"tiktok/config"
	"time"
)

// sha256加密
func SHA256(password, salt string) string {
	code := []byte(salt)
	saltedPassword := append([]byte(password), ':')
	saltedPassword = append(saltedPassword, code...)

	hash := sha256.Sum256(saltedPassword)
	return hex.EncodeToString(hash[:])
}

// 用于解析token
type UserClaim struct {
	Id uint
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id uint) string {
	exp := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	uc := UserClaim{
		Id:             id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: exp},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	var JwtKey = []byte(config.Conf.GetString("auth.jwt_key"))
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

// 生成随机salt
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func GenerateSalt() string {
	b := make([]byte, 8)
	for i := 0; i < 8; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

// 检验token
func VerifyToken(id uint, tokenString string) bool {
	var JwtKey = []byte(config.Conf.GetString("auth.jwt_key"))
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return false
	}
	// Check if the token is valid
	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		if claims.Id != id {
			return false
		}
		return true
	} else {
		return false
	}
}
