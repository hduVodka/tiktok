package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserId   uint
	ToUserId uint
	Content  string `gorm:"not null"`
}
