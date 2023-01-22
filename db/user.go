package db

import (
	"errors"
	"tiktok/models"
	"tiktok/utils"
)

func InsertNewUser(user *models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return errors.New("插入用户数据失败")
	}
	return nil
}

func SearchUser(user *models.User) bool {
	var userExist models.User

	// 通过用户名找到salt
	if err := db.Where("username = ?", user.Username).First(&userExist).Error; err != nil {
		return false
	}
	if db.Where("username = ? AND password = ?", user.Username, utils.SHA256(user.Password, userExist.Salt)).First(&user).Error != nil {
		return false
	}
	user.ID = userExist.ID
	return true
}

func FindUserInfo(userId uint, user *models.User) bool {
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}
