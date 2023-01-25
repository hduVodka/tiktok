package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint
	User    *User `gorm:"foreignKey:UserID"`
	VideoID uint
	Video   *Video `gorm:"foreignKey:VideoID"`
}
