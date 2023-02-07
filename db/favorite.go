package db

import (
	"context"
	"tiktok/log"
	"tiktok/models"
)

const favoriteVideoKey = "fvk"

//const favoriteCountKey = "fck"

// 用 |作分隔符
const separator = "|"

func AddFavorite(ctx context.Context, f *models.Favorite) error {
	res := db.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).FirstOrCreate(&f)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return nil
	}
	return IncreaseVideoFavoriteCount(ctx, f.VideoID, 1)
}

func RemoveFavorite(ctx context.Context, userId, videoId uint) error {
	res := db.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&models.Favorite{})
	if res.Error != nil {
		log.Errorln("delete favorite error: ", res.Error)
		return ErrDatabase
	}
	if res.RowsAffected == 0 {
		return nil
	}
	return IncreaseVideoFavoriteCount(ctx, videoId, -1)
}

func GetFavoriteListByUserID(userId uint) ([]models.Video, error) {
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
