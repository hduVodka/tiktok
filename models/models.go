package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
)

var db *gorm.DB

func Init() {
	source := config.Conf.GetString("server.mysql_source")
	database, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to connect mysql:%v", err)
	}
	db = database
}
