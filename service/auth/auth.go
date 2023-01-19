package auth

import (
	"regexp"
	"tiktok/models"
	"tiktok/utils"
)

// 密码加盐生成
func Encrypt(user *models.User) *models.User {
	user.Salt = utils.GenerateSalt()
	hashedPassword := utils.SHA256(user.Password, user.Salt)
	user.Password = hashedPassword
	return user
}

// 根据已有盐和明文密码生成加密密码
func EncryptPassword(user *models.User, salt string) *models.User {
	user.Salt = salt
	hashedPassword := utils.SHA256(user.Password, user.Salt)
	user.Password = hashedPassword
	return user
}

// 验证用户名是否合法
func CheckLegal(user *models.User) bool {
	if len(user.Password) < 6 || len(user.Password) >= 32 || len(user.Username) >= 32 {
		return false
	}
	emailPattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(emailPattern)
	return reg.MatchString(user.Username)
}
