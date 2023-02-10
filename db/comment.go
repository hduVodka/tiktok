package db

import (
	"tiktok/log"
	"tiktok/models"
	"time"
)

// InsertComment 向表中写入评论，如果已经存在，就更新评论信息
func InsertComment(c *models.Comment) (int64, error) {
	var count int64
	if err := db.Model(&models.Comment{}).Where("user_id = ? AND video_id = ?", c.UserID, c.VideoID).Count(&count).Error; err != nil {
		return count, err
	}
	if count > 0 {
		err := UpdateComment(c)
		if err != nil {
			return count, err
		}
		return count, nil
	}
	if err := db.Create(c).Error; err != nil {
		log.Error("insert comment error:", err)
		return count, err
	}
	return count, nil
}

// UpdateComment 更新评论信息
func UpdateComment(c *models.Comment) error {
	if err := db.Model(&models.Comment{}).Where("video_id = ? AND user_id = ?", c.VideoID, c.UserID).Updates(models.Comment{
		Content:    c.Content,
		CreateTime: time.Now().Unix(),
	}).Error; err != nil {
		log.Error("update comment error:", err)
		return err
	}
	return nil
}

// DeleteComment 从表中删除评论
func DeleteComment(c *models.Comment) error {
	var count int64
	if err := db.Model(&models.Comment{}).Where("user_id = ? AND video_id = ?", c.UserID, c.VideoID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		log.Error("delete comment error")
		return nil
	}
	if err := db.Delete(c, "user_id = ? AND video_id = ?", c.UserID, c.VideoID).Error; err != nil {
		log.Error("delete comment error", err)
		return err
	}
	return nil
}

// GetCommentListByVideoId 获取评论列表
func GetCommentListByVideoId(videoId uint) ([]models.Comment, error) {
	comments := make([]models.Comment, 0)
	if err := db.Model(&models.Comment{}).
		Where("video_id = ? AND deleted_at IS NULL", videoId).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		log.Error("get comment list error ", err)
		return nil, err
	}
	return comments, nil
}

// GetUserMessage 获取关联User信息
func GetUserMessage(v models.Comment) models.User {
	user := models.User{}
	db.Model(&v).Association("User").Find(&user)
	return user
}

// GetVideoMessage 获取关联video信息
func GetVideoMessage(v models.Comment) models.Video {
	video := models.Video{}
	db.Model(&v).Association("Video").Find(&video)
	return video
}

// CountCommentById 统计评论数
func CountCommentById(VideoID uint) (int64, error) {
	var count int64
	if err := db.Model(&models.Comment{}).Where("video_id = ?", VideoID).Count(&count).Error; err != nil {
		log.Error("count comment error:", err)
		return -1, err
	}
	return count, nil
}
