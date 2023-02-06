package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
	"time"
)

const FeedPageSize = 30

func GetFeedByTime(ctx context.Context, t time.Time) ([]models.Video, error) {

	// get ids from cache
	strIds, err := rdb.Do(ctx, "ZRANGE", "video:feed", t.UnixMilli(), 0, "BYSCORE", "REV", "limit", 0, FeedPageSize).StringSlice()
	if err != nil {
		log.Errorf("get feed ids from cache fail:%v", err)
		return nil, ErrDatabase
	}

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
	res := db.Where("id in ?", toGetFromDB).Find(&dbVd)
	if res.Error != nil {
		log.Errorf("get feed from db fail:%v", res.Error)
		return nil, ErrDatabase
	}

	// set cache
	for _, v := range dbVd {
		if err := rdb.HSet(ctx, fmt.Sprintf("video:%d", v.ID), v).Err(); err != nil {
			log.Errorf("set cache fail: %v", err)
		}
		if err := rdb.Expire(ctx, fmt.Sprintf("video:%d", v.ID), 3*time.Minute).Err(); err != nil {
			log.Error("set expire fail: %v", err)
		}
	}

	return append(videos, dbVd...), nil
}

func InsertVideo(ctx context.Context, video *models.Video) error {
	res := db.Save(video)
	if res.Error != nil {
		log.Errorf("insert video fail:%v", res.Error)
		return ErrDatabase
	}
	if err := rdb.ZAdd(ctx, "video:feed", redis.Z{
		Score:  float64(video.CreatedAt.UnixMilli()),
		Member: video.ID,
	}).Err(); err != nil {
		log.Errorf("insert video to feed cache fail:%v", err)
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
