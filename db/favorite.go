package db

import (
	"tiktok/log"
	"tiktok/models"
)

func InsertFavorite(f *models.Favorite) error {
	if err := db.Create(f).Error; err != nil {
		log.Error("insert favorite error: ", err)
		return err
	}
	return nil
}

func DeleteFavorite(f *models.Favorite) error {
	if err := db.Delete(f, "user_id = ? AND video_id = ? ", f.UserID, f.VideoID).Error; err != nil {
		log.Error("delete favorite error: ", err)
		return err
	}
	return nil
}

func GetFavoriteListByUserID(f *models.Favorite) ([]models.Video, error) {
	var videos []models.Video
	if err := db.Model(&models.Video{}).Joins("join favorites AS f on f.video_id = videos.id").Where("f.user_id = ? AND f.deleted_at IS NULL", f.UserID).Order("f.created_at desc").Find(&videos).Error; err != nil {
		log.Error("get favorite list error: ", err)
		return nil, err
	}

	return videos, nil
}

func IsFavorite(f *models.Favorite) bool {
	var count int64
	if err := db.Model(f).Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Count(&count).Error; err != nil {
		log.Error("count favorite error: ", err)
		return false
	}
	return count > 0
}

func CountFavoriteByID(videoID uint) (int64, error) {
	var count int64
	if err := db.Model(&models.Favorite{}).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
		log.Error("count favorite error: ", err)
		return 0, err
	}
	return count, nil
}
