package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"tiktok/config"
	"tiktok/models"
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
	Id       int
	Password string
	Username string
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user *models.User) string {
	uc := UserClaim{
		Password: user.Password,
		Username: user.Username,
	}
	// 用jwt中的方法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 对token进行加密,key在config里
	var JwtKey = []byte(config.Conf.GetString("auth.jwt_key"))
	tokenString, _ := token.SignedString([]byte(JwtKey))
	return tokenString
}

// 生成随机salt
func GenerateSalt() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
