package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID            uint           `gorm:"primarykey" redis:"ID"`
	CreatedAt     time.Time      `redis:"CreatedAt"`
	UpdatedAt     time.Time      `redis:"UpdatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorId      uint           `redis:"AuthorId"`
	Title         string         `redis:"Title"`
	PlayUrl       string         `redis:"PlayUrl"`
	CoverUrl      string         `redis:"CoverUrl"`
	CommentCount  uint           `redis:"CommentCount"`
	FavoriteCount uint           `redis:"FavoriteCount"`
}
