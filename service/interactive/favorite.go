package interactive

import (
	"context"
	"errors"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
)

const AddFavorite = 1
const CancelFavorite = 2

func FavoriteAction(ctx context.Context, favorite *models.Favorite, actionType int) error {
	// 执行点赞或取消点赞操作
	if actionType == AddFavorite {
		return db.AddFavorite(ctx, favorite)
	}
	if actionType == CancelFavorite {
		return db.RemoveFavorite(ctx, favorite.UserID, favorite.VideoID)
	}
	return errors.New("请求参数错误")
}

func FavoriteList(ctx context.Context, userId uint) ([]dto.Video, error) {
	videoList, err := db.GetFavoriteListByUserID(userId)
	if err != nil {
		return nil, err
	}
	return dto.FromVideoModels(ctx, userId, videoList)

}
