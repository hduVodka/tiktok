package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiktok/log"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId uint
	Title    string
	PlayUrl  string
	CoverUrl string
}

func GetFeedByTime(t time.Time) ([]Video, error) {
	var list []Video
	res := db.Limit(30).Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "created_at",
		},
		Desc: true,
	}).Find(&list)
	if res.Error != nil {
		log.Errorf("get feed fail:%v", res.Error)
		return nil, ErrDatabase
	}
	return list, nil
}

func InsertVideo(video *Video) error {
	res := db.Save(video)
	if res.Error != nil {
		log.Errorf("insert video fail:%v", res.Error)
		return ErrDatabase
	}
	return nil
}
