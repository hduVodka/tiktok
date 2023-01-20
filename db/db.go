package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
	"tiktok/log"
	"tiktok/models"
)

var db *gorm.DB

func Init() {
	ModelInit()
	//todo: init redis
}

func ModelInit() {
	source := config.Conf.GetString("server.mysql_source")
	database, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to connect mysql:%v", err)
	}
	db = database
	if err := db.AutoMigrate(&models.User{}, &models.Video{}, &models.Favorite{}, &models.Comment{}, &models.Follow{}); err != nil {
		log.Fatalf("fail to migrate models:%v", err)
	}
}
