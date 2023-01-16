package helper

import (
	"github.com/dgrijalva/jwt-go"
	"tiktok/config"
)

// 用于解析token
type UserClaim struct {
	Id       int
	Password string
	Username string
	jwt.StandardClaims
}

func GenerateToken(Password, Username string) (string, error) {
	uc := UserClaim{
		Password: Password,
		Username: Username,
	}
	// 用jwt中的方法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 对token进行加密,key在config里
	var JwtKey = []byte(config.Conf.GetString("JwtKey"))
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
