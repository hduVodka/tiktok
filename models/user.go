package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64  `form:"id" json:"id"`
	Nickname string `form:"nickname" json:"nickname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Salt     string
}
