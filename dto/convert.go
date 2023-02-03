package dto

import (
	"context"
	"tiktok/db"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
)

func FromVideoModels(ctx context.Context, userId uint, videos []models.Video) ([]Video, error) {
	res := make([]Video, len(videos))
	for i, v := range videos {
		author := db.GetUser(v.AuthorId)
		isFav, err := db.IsFavorite(userId, v.ID)
		if err != nil {
			log.Errorf("fail to check favorite:%v", err)
		}
		res[i] = Video{
			Id: v.ID,
			Author: User{
				Id:            author.ID,
				Name:          author.Username,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       utils.GenUrl(ctx, v.PlayUrl),
			CoverUrl:      utils.GenUrl(ctx, v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFav,
			Title:         v.Title,
		}
	}
	return res, nil
}

func FromMessageModels(models []models.Message) []Message {
	messageList := make([]Message, len(models))
	for i, v := range models {
		messageList[i] = Message{
			Id:         v.ID,
			Content:    v.Content,
			FromUserId: v.UserId,
			ToUserId:   v.ToUserId,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return messageList
}
