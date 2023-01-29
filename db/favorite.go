package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"tiktok/log"
	"tiktok/models"
)

const favoriteVideoKey = "fvk"
const favoriteCountKey = "fck"

// 用 |作分隔符
const separator = "|"

func InsertFavorite(f *models.Favorite) error {
	res, err := rdb.HGet(context.Background(), favoriteVideoKey,
		fmt.Sprintf("%d%s%d", f.UserID, separator, f.VideoID)).Result()
	// 字段不存在或者字段值为0
	if err != nil || res == "0" {
		if err := rdb.HSet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", f.UserID, separator, f.VideoID), 1).Err(); err != nil {
			log.Error("set cache error: ", err)
			return err
		}
		//if err = rdb.HIncrBy(context.Background(),
		//	favoriteCountKey,strconv.FormatUint(uint64(f.VideoID),10),
		//	1).Err();err != nil{
		//	log.Error("incr count error: ", err)
		//	return err
		//}
	}

	return nil
}

func DeleteFavorite(userId, videoId uint) error {
	res, err := rdb.HGet(context.Background(), favoriteVideoKey,
		fmt.Sprintf("%d%s%d", userId, separator, videoId)).Result()
	// 字段存在且字段值为1
	if err == nil && res == "1" {
		if err := rdb.HSet(context.Background(), favoriteVideoKey,
			fmt.Sprintf("%d%s%d", userId, separator, videoId), 0).Err(); err != nil {
			log.Error("set cache error: ", err)
			return err
		}
		//if err = rdb.HIncrBy(context.Background(),
		//	favoriteCountKey,strconv.FormatUint(uint64(videoId),10),
		//	-1).Err();err != nil{
		//	log.Error("incr count error: ", err)
		//	return err
		//}
	}

	return nil
}

func GetFavoriteListByUserID(userId uint) ([]models.Video, error) {
	var mutex sync.Mutex
	mutex.Lock()
	err := updateDB()
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	err = deleteCache()
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	mutex.Unlock()

	var videos []models.Video
	if err := db.Model(&models.Video{}).
		Joins("join favorites AS f on f.video_id = videos.id").
		Where("f.user_id = ? AND f.deleted_at IS NULL", userId).
		Order("f.created_at desc").
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

//func CountFavoriteByID(videoId uint) (int64, error) {
//	var res string
//	var err error
//	if res,err = rdb.HGet(context.Background(),favoriteCountKey, strconv.FormatUint(uint64(videoId), 10)).
//		Result();err != nil{
//			log.Error("get count error: ", err)
//			return 0,err
//	}
//	return strconv.ParseInt(res,10,64)
//}

func updateDB() error {
	result, err := rdb.HGetAll(context.Background(), favoriteVideoKey).Result()
	if err != nil {
		log.Error("get cache error: ", err)
		return err
	}

	var userId, videoId uint

	for k, v := range result {
		// 将redis中的字符串解析成userId和videoId
		str := strings.Split(k, separator)
		id, _ := strconv.Atoi(str[0])
		userId = uint(id)
		id, _ = strconv.Atoi(str[1])
		videoId = uint(id)
		if v == "1" {
			exist, err := IsFavorite(userId, videoId)
			if err != nil {
				return err
			}
			if !exist {
				if err := db.Create(&models.Favorite{
					UserID:  userId,
					VideoID: videoId,
				}).Error; err != nil {
					log.Error("insert favorite error: ", err)
					return err
				}
			}
		} else {
			if err := db.Delete(&models.Favorite{},
				"user_id = ? AND video_id = ? ", userId, videoId).Error; err != nil {
				log.Error("delete favorite error: ", err)
				return err
			}
		}
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
