package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
	"time"
)

func GetFeedByTime(ctx context.Context, t time.Time) ([]models.Video, error) {

	// get ids from cache
	strIds, err := rdb.Do(ctx, "ZRANGE", "video:feed", t.UnixMilli(), 0, "BYSCORE", "REV", "limit", 0, 30).StringSlice()
	if err != nil {
		log.Errorf("get feed from cache fail:%v", err)
		return nil, ErrDatabase
	}

	//todo: when out of cache, read from db

	// get video from cache
	var toGetFromDB []string
	var videos []models.Video
	for _, v := range strIds {
		mp, err := rdb.HGetAll(ctx, "video:"+v).Result()
		if err != nil {
			log.Errorf("get video from cache fail:%v", err)
			continue
		}
		vd := utils.Scan[models.Video](mp)
		if vd.ID == 0 {
			toGetFromDB = append(toGetFromDB, v)
			continue
		}
		videos = append(videos, *vd)
	}

	if len(toGetFromDB) == 0 {
		return videos, nil
	}

	// get uncached video from db
	var dbVd []models.Video
	res := db.Where("id in ?", toGetFromDB).
		Limit(30).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "created_at",
			},
			Desc: true,
		}).Find(&dbVd)
	if res.Error != nil {
		log.Errorf("get feed fail:%v", res.Error)
		return nil, ErrDatabase
	}

	// set cache
	for _, v := range dbVd {
		if err := rdb.HSet(ctx, fmt.Sprintf("video:%d", v.ID), v).Err(); err != nil {
			log.Error(err)
		}
	}

	return append(videos, dbVd...), nil
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
	res := db.Where("author_id = ?", id).Find(&list)
	if res.Error != nil {
		log.Errorf("get video list fail:%v", res.Error)
		return nil, ErrDatabase
	}
	return list, nil
}

func IncreaseVideoFavoriteCount(id uint, count int) error {
	res := db.Model(&models.Video{}).Where("id=?", id).Update("favorite_count", gorm.Expr("favorite_count + ?", count))
	if res.Error != nil {
		log.Errorf("increase video favorite count fail:%v", res.Error)
		return ErrDatabase
	}
	return nil
}

func IncreaseVideoCommentCount(id uint, count int) error {
	res := db.Model(&models.Video{}).Where("id=?", id).Update("comment_count", gorm.Expr("comment_count + ?", count))
	if res.Error != nil {
		log.Errorf("increase video comment count fail:%v", res.Error)
		return ErrDatabase
	}
	return nil
}
