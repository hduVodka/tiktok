package message

import (
	"context"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
)

func GetList(ctx context.Context, userId uint, toUserId uint) ([]dto.Message, error) {
	message, err := db.SearchMessage(ctx, userId, toUserId)
	if err != nil {
		return nil, err
	}
	return dto.FromMessageModels(message), nil
}

func Send(ctx context.Context, userId, toUserId uint, content string) error {
	return db.InsertMessage(ctx, &models.Message{
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
	})
}
