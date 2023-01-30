package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID     uint
	User       *User `gorm:"foreignKey:UserID"`
	VideoID    uint
	Video      *Video `gorm:"foreignKey:VideoID"`
	Content    string `gorm:"not null"`
	CreateTime int64  `gorm:"autoCreateTime"`
}
