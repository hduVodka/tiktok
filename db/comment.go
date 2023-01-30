package db

import (
	"tiktok/log"
	"tiktok/models"
	"time"
)

// InsertComment 向表中写入评论，如果已经存在，就更新评论信息
func InsertComment(c *models.Comment) error {
	var count int64
	if err := db.Model(&models.Comment{}).Where("user_id = ? AND video_id = ?", c.UserID, c.VideoID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := UpdateComment(c)
		if err != nil {
			return err
		}
	}
	if err := db.Create(c).Error; err != nil {
		log.Error("insert comment error:", err)
		return err
	}
	return nil
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
	var comments []models.Comment
	if err := db.Raw("SELECT v.* FROM comments AS v WHERE v.video_id = ? AND v.deleted_at IS NULL", videoId).Scan(&comments).Error; err != nil {
		log.Error("get comments list error:", err)
		return nil, err
	}
	return comments, nil
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
