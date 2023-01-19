package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

func TestGetFeedByTime(t *testing.T) {
	db, _ = gorm.Open(mysql.Open(os.Getenv("mysql_source")), &gorm.Config{})
	feed, err := GetFeedByTime(time.Now())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(feed)
}
