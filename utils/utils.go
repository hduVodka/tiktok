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

//todo: 修复下面的函数，jwt请使用jwt的方式校验

// 检验token
func CheckToken(userId int64, tokenString string) bool {
	/*
		var userExist models.User
		db := db.ModelInit()
		if err := db.Where("id = ?", userId).First(&userExist).Error; err != nil {
			return false
		}
		tokenRight := GenerateToken(&userExist)
		if tokenRight != tokenString {
			return false
		}
	*/
	return true
}
