package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId uint
	Title    string
	PlayUrl  string
	CoverUrl string
}
