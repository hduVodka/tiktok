package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
	"tiktok/log"
)

var db *gorm.DB

func Init() {
	source := config.Conf.GetString("server.mysql_source")
	database, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to connect mysql:%v", err)
	}
	db = database
	if err := db.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Comment{}, &Follow{}); err != nil {
		log.Fatalf("fail to migrate models:%v", err)
	}
}
