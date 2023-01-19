package models

import (
	"gorm.io/gorm"
	"tiktok/log"
)

type User struct {
	gorm.Model
	ID            int64  `form:"id" json:"id"`
	Nickname      string `form:"nickname" json:"nickname"`
	Username      string `form:"username" json:"username"`
	Password      string `form:"password" json:"password"`
	FollowerCount int64  `form:"follower_count" json:"follower_count"`
	FollowCount   int64  `form:"follow_count" json:"follow_count"`
	IsFollow      bool   `form:"is_follow" json:"is_follow"`
	Salt          string
}

// 真实插入前要对密码进行加密
func InsertUser(user *User) error {
	user.FollowCount = 0
	user.FollowerCount = 0
	user.IsFollow = false
	res := db.Save(user)
	if res.Error != nil {
		log.Errorf("insert user fail:%v", res.Error)
		return ErrDatabase
	}
	return nil
}

// 检查用户是否存在
func CheckUsername(user *User) error {
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return ErrDatabase
		}
	} else {
		return ErrDatabase
	}
	return nil
}

// 根据用户名查找ID
func SearchUser(user *User) (*User, error) {
	var userExist User
	// 通过用户名找到salt
	if err := db.Where("username = ?", user.Username).First(&userExist).Error; err != nil {
		return &userExist, ErrUserNotFound
	}
	user.ID = userExist.ID
	return &userExist, nil
}
