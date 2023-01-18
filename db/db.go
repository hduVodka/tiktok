package db

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/models"
	"tiktok/utils"
)

func Init() {
	//todo: init redis

	models.Init()
}

func CheckUsername(user *models.User) bool {
	db := models.Init()
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false
		}
	} else {
		return false
	}
	return true
}

func InsertNewUser(user *models.User) error {
	// 清除意外数据
	user.FollowerCount = 0
	user.FollowCount = 0
	user.IsFollow = false
	db := models.Init()
	if err := db.Create(&user).Error; err != nil {
		return errors.New("插入用户数据失败")
	}
	return nil
}

func SearchUser(user *models.User) bool {
	db := models.Init()
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
