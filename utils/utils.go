package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"tiktok/config"
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
	uc := UserClaim{
		Id: id,
	}
	// 用jwt中的方法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 对token进行加密,key在config里
	var JwtKey = []byte(config.Conf.GetString("auth.jwt_key"))
	tokenString, _ := token.SignedString([]byte(JwtKey))
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
func CheckToken(userID uint, tokenString string) bool {
	if tokenString == GenerateToken(userID) {
		return true
	}
	return false
}
