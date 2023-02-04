package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
	"tiktok/log"
	"tiktok/models"
)

const favoriteVideoKey = "fvk"

//const favoriteCountKey = "fck"

// 用 |作分隔符
const separator = "|"

func AddFavorite(f *models.Favorite) error {
	/*
		res, err := rdb.HGet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", f.UserID, separator, f.VideoID)).Result()

		if err == nil && res == "1" {
			return errors.New("已经点赞过了")
		}

		// 字段不存在或者字段值为0
		if err := rdb.HSet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", f.UserID, separator, f.VideoID), 1).Err(); err != nil {
			log.Error("set cache error: ", err)
			return err
		}

		return nil
	*/
	if err := db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).FirstOrCreate(&f)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return nil
		}
		return IncreaseVideoFavoriteCount(f.VideoID, 1)
	}); err != nil {
		log.Errorln("create favorite error: ", err)
		return ErrDatabase
	}
	return nil
}

func RemoveFavorite(userId, videoId uint) error {
	/*
		res, err := rdb.HGet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", userId, separator, videoId)).Result()

		if err != nil || res == "0" {
			return errors.New("还没有点赞过")
		}

		// 字段存在且字段值为1
		if err := rdb.HSet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", userId, separator, videoId), 0).Err(); err != nil {
			log.Error("set cache error: ", err)
			return err
		}

		return nil
	*/
	if err := db.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&models.Favorite{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Errorln("delete favorite error: ", err)
		return ErrDatabase
	}
	return IncreaseVideoFavoriteCount(videoId, -1)
}

func GetFavoriteListByUserID(userId uint) ([]models.Video, error) {

	//if err := updateDB(); err != nil {
	//	return nil, err
	//}

	var videos []models.Video
	if err := db.Where("id in (?)",
		db.Model(&models.Favorite{}).
			Select("video_id").
			Where("user_id = ?", userId)).
		Find(&videos).Error; err != nil {
		log.Error("get favorite list error: ", err)
		return nil, err
	}

	return videos, nil
}

func IsFavorite(userId, videoId uint) (bool, error) {
	var count int64
	if err := db.Model(&models.Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&count).Error; err != nil {
		log.Error("count favorite error: ", err)
		return false, err
	}
	return count > 0, nil
}

func updateDB() error {
	var mutex sync.Mutex
	mutex.Lock()
	err := insertMysql()
	if err != nil {
		mutex.Unlock()
		return err
	}
	err = deleteCache()
	if err != nil {
		mutex.Unlock()
		return err
	}
	mutex.Unlock()
	return nil
}

func insertMysql() error {
	result, err := rdb.HGetAll(context.Background(), favoriteVideoKey).Result()
	if err != nil {
		log.Error("get cache error: ", err)
		return err
	}

	var userId, videoId uint

	for k, v := range result {
		// 将redis中的字符串解析成userId和videoId
		userId, videoId = parseFavoriteKey(k)
		if v == "1" {
			// 点赞
			if err := insertFavorite(userId, videoId); err != nil {
				return err
			}
		} else {
			// 取消点赞
			if err := deleteFavorite(userId, videoId); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseFavoriteKey(key string) (uint, uint) {
	str := strings.Split(key, separator)
	id, _ := strconv.Atoi(str[0])
	userId := uint(id)
	id, _ = strconv.Atoi(str[1])
	videoId := uint(id)
	return userId, videoId
}

func insertFavorite(userId, videoId uint) error {
	exist, err := IsFavorite(userId, videoId)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	if err := db.Create(&models.Favorite{
		UserID:  userId,
		VideoID: videoId,
	}).Error; err != nil {
		log.Error("insert favorite error: ", err)
	}
	return nil
}

func deleteFavorite(userId, videoId uint) error {
	if err := db.Delete(&models.Favorite{},
		"user_id = ? AND video_id = ? ", userId, videoId).Error; err != nil {
		log.Error("delete favorite error: ", err)
	}
	return nil
}

func deleteCache() error {
	if err := rdb.Del(context.Background(), favoriteVideoKey).Err(); err != nil {
		log.Error("delete cache error: ", err)
		return err
	}
	return nil
}
