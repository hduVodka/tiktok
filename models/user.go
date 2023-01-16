package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64
	Nickname string
	Username string
	Password string
}
