package db

import (
	"gorm.io/gorm/clause"
	"tiktok/log"
	"tiktok/models"
	"time"
)

func GetFeedByTime(t time.Time) ([]models.Video, error) {
	var list []models.Video
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

func InsertVideo(video *models.Video) error {
	res := db.Save(video)
	if res.Error != nil {
		log.Errorf("insert video fail:%v", res.Error)
		return ErrDatabase
	}
	return nil
}

func GetVideoListById(id uint) ([]models.Video, error) {
	var list []models.Video
	res := db.Where("author_id=?", id).Find(&list)
	if res.Error != nil {
		log.Errorf("get video list fail:%v", res.Error)
		return nil, ErrDatabase
	}
	return list, nil
}
