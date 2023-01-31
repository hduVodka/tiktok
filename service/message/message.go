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
	return arrayModel2Dto(message), nil
}

func Send(ctx context.Context, userId, toUserId uint, content string) error {
	return db.InsertMessage(ctx, &models.Message{
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
	})
}

func arrayModel2Dto(models []models.Message) []dto.Message {
	messageList := make([]dto.Message, len(models))
	for i, v := range models {
		messageList[i] = dto.Message{
			Id:         v.ID,
			Content:    v.Content,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return messageList
}
