package db

import (
	"gorm.io/gorm"
	"tiktok/models"
)

var db *gorm.DB

func Init() {
	//todo: init redis

	db = models.Init()
}
