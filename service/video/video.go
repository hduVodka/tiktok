package video

import (
	"context"
	"tiktok/dto"
	"tiktok/models"
	"time"
)

func GetFeed(ctx context.Context, latestTime time.Time) ([]dto.Video, time.Time, error) {
	videos, err := models.GetFeedByTime(latestTime)
	if err != nil {
		return nil, time.Now(), err
	}

	oldest := time.Now()
	res := make([]dto.Video, len(videos))
	for i, v := range videos {
		// todo cache favorite and comment count
		res[i] = dto.Video{
			Id:            v.ID,
			Author:        dto.User{},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    0,
			Title:         v.Title,
		}
		if i == len(videos)-1 {
			oldest = v.CreatedAt
		}
	}
	return nil, oldest, nil
}
