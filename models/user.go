package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname      string `form:"nickname" json:"nickname"`
	Username      string `form:"username" json:"username"`
	Password      string `form:"password" json:"password"`
	FollowerCount uint   `form:"follower_count" json:"follower_count"`
	FollowCount   uint   `form:"follow_count" json:"follow_count"`
	Salt          string
}
