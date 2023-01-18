package models

import (
	"gorm.io/gorm"
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
