package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tiktok/log"
	"tiktok/models"
)

func InsertMessage(ctx context.Context, message *models.Message) error {
	if err := db.Create(message).Error; err != nil {
		log.Errorln(err)
		return ErrDatabase
	}
	return nil
}

func SearchMessage(ctx context.Context, userId, toUserId uint) ([]models.Message, error) {
	var messageList []models.Message
	if err := db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&messageList).Error; err != nil {
		log.Errorln(err)
		return nil, ErrDatabase
	}
	return messageList, nil
}

func LatestMessageBetween(ctx context.Context, userId1, userId2 uint) (*models.Message, error) {
	var message models.Message
	if err := db.Where(
		db.Where("user_id = ?", userId1).Where("to_user_id = ?", userId2),
	).Or(
		db.Where("user_id = ?", userId2).Where("to_user_id = ?", userId1),
	).Order("created_at desc").First(&message).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorln(err)
			return nil, ErrDatabase
		}
	}
	return &message, nil
}
