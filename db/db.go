package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
	"tiktok/log"
	"tiktok/models"
)

var db *gorm.DB
var rdb *redis.Client

func Init() {
	ModelInit()
	RedisInit()
}

func ModelInit() {
	source := config.Conf.GetString("server.mysql_source")
	database, err := gorm.Open(mysql.Open(source), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("fail to connect mysql:%v", err)
	}
	db = database
	if err := db.AutoMigrate(&models.User{}, &models.Video{}, &models.Favorite{}, &models.Comment{}, &models.Follow{}); err != nil {
		log.Fatalf("fail to migrate models:%v", err)
	}
}

func RedisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Conf.GetString("server.redis.addr"),
		Password: config.Conf.GetString("server.redis.password"),
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("fail to connect redis:%v", err)
	}
}
