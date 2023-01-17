package auth

import (
	"tiktok/models"
	"tiktok/utils"
)

func Encrypt(user *models.User) {
	hashedPassword := utils.SHA256(user.Password)
	user.Password = string(hashedPassword)
	println(user.Password)
}
