package db

import (
	"context"
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
