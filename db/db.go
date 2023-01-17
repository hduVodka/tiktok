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

func CheckUsername(user *models.User) (bool, error) {
	db := models.Init()
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, errors.New("连接数据库失败")
		}
	} else {
		return false, errors.New("用户已存在")
	}
	return true, nil
}

func InsertNewUser(user *models.User) error {
	db := models.Init()
	if err := db.Create(&user).Error; err != nil {
		return errors.New("插入用户数据失败")
	}
	return nil
}

func SearchUser(user *models.User) (bool, error) {
	db := models.Init()
	var userExist models.User
	if db.Where("username = ? AND password = ?", user.Username, utils.SHA256(user.Password)).First(&userExist).Error != nil {
		return false, errors.New("用户不存在或密码错误")
	}
	return true, nil
}
