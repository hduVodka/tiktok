package models

import (
	"gorm.io/gorm"
	"tiktok/log"
)

type Favorite struct {
	gorm.Model
	UserID  uint
	User    *User `gorm:"foreignKey:UserID"`
	VideoID uint
	Video   *Video `gorm:"foreignKey:VideoID"`
}

// TODO: redis 缓存

func (f *Favorite) InsertFavorite() error {
	if err := db.Create(f).Error; err != nil {
		log.Error("insert favorite error: ", err)
		return err
	}
	return nil
}

func (f *Favorite) DeleteFavorite() error {
	if err := db.Delete(f, "user_id = ? AND video_id = ? ", f.UserID, f.VideoID).Error; err != nil {
		log.Error("delete favorite error: ", err)
		return err
	}
	return nil
}

func (f *Favorite) GetFavoriteListByUserID() ([]Video, error) {
	var videos []Video
	if err := db.Raw("SELECT v.* FROM videos AS v JOIN favorites AS f ON v.id = f.video_id WHERE f.user_id = ? AND f.deleted_at IS NULL ORDER BY f.created_at desc", f.UserID).Scan(&videos).Error; err != nil {
		log.Error("get favorite list error: ", err)
		return nil, err
	}
	return videos, nil
}

func (f *Favorite) IsFavorite() bool {
	var count int64
	if err := db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Count(&count).Error; err != nil {
		log.Error("count favorite error: ", err)
		return false
	}
	return count > 0
}

func CountFavoriteByID(videoID uint) (int64, error) {
	var count int64
	if err := db.Model(&Favorite{}).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
		log.Error("count favorite error: ", err)
		return 0, err
	}
	return count, nil
}
