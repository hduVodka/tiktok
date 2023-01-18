package models

import (
	"gorm.io/gorm"
	"tiktok/log"
	"time"
)

type Comment struct {
	gorm.Model
	UserID     uint
	User       *User `gorm:"foreignKey:UserID"`
	VideoID    uint
	Video      *Video `gorm:"foreignKey:VideoID"`
	Content    string `gorm:"not null"`
	CreateTime int64  `gorm:"autoCreateTime"`
}

// InsertComment 向表中写入评论，如果已经存在，就更新评论信息
func (c *Comment) InsertComment() error {
	if c.IsInSQL() {
		err := c.UpdateComment()
		return err
	} else {
		if err := db.Create(c).Error; err != nil {
			log.Error("insert comment error:", err)
			return err
		}
	}
	return nil
}

// UpdateComment 更新评论信息
func (c *Comment) UpdateComment() error {
	if err := db.Model(&Comment{}).Where("video_id = ? AND user_id = ?", c.VideoID, c.UserID).Updates(Comment{
		Content:    c.Content,
		CreateTime: time.Now().Unix(),
	}).Error; err != nil {
		log.Error("update comment error:", err)
		return err
	}
	return nil
}

// IsInSQL 判断数据是否在表中
func (c *Comment) IsInSQL() bool {
	var count int64
	if err := db.Model(&Comment{}).Where("user_id = ? AND video_id = ?", c.UserID, c.VideoID).Count(&count).Error; err != nil {
		log.Error("count comment error:", err)
		return false
	}
	return count > 0
}

// DeleteComment 从表中删除评论
func (c *Comment) DeleteComment() error {
	if err := db.Delete(c, "user_id = ? AND video_id = ?", c.UserID, c.VideoID).Error; err != nil {
		log.Error("delete comment error", err)
		return err
	}
	return nil
}

// GetCommentListByVideoId 获取评论列表
func GetCommentListByVideoId(videoId uint) ([]*Comment, error) {
	var comments []*Comment
	if err := db.Raw("SELECT v.* FROM comments AS v WHERE v.video_id = ? AND v.deleted_at IS NULL", videoId).Scan(&comments).Error; err != nil {
		log.Error("get comments list error:", err)
		return nil, err
	}
	return comments, nil
}

// CountCommentById 统计评论数
func CountCommentById(VideoID uint) (int64, error) {
	var count int64
	if err := db.Model(&Comment{}).Where("video_id = ?", VideoID).Count(&count).Error; err != nil {
		log.Error("count comment error:", err)
		return -1, err
	}
	return count, nil
}
