package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"tiktok/config"
	"tiktok/models"
)

// sha256加密
func SHA256(password string) string {
	salt := []byte(config.Conf.GetString("auth.salt"))
	saltedPassword := append([]byte(password), ':')
	saltedPassword = append(saltedPassword, salt...)

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

func GenerateToken(user *models.User) (string, error) {
	uc := UserClaim{
		Password: user.Password,
		Username: user.Username,
	}
	// 用jwt中的方法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 对token进行加密,key在config里
	var JwtKey = []byte(config.Conf.GetString("auth.jwt_key"))
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
