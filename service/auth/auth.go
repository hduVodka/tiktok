package auth

import (
	"regexp"
	"tiktok/models"
	"tiktok/utils"
)

func Encrypt(user *models.User) *models.User {
	user.Salt = utils.GenerateSalt()
	hashedPassword := utils.SHA256(user.Password, user.Salt)
	user.Password = string(hashedPassword)
	return user
}

// 验证用户名是否合法
func CheckLegal(user *models.User) bool {
	if len(user.Password) < 6 {
		return false
	}
	emailPattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(emailPattern)
	return reg.MatchString(user.Username)
}
