package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserId   uint `gorm:"uniqueIndex:idx_user_id_to_user_id"`
	ToUserId uint `gorm:"uniqueIndex:idx_user_id_to_user_id"`
}
