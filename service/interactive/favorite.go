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

func FavoriteAction(favorite *models.Favorite, actionType int) error {
	// 执行点赞或取消点赞操作
	if actionType == AddFavorite {
		if err := db.AddFavorite(favorite); err != nil {
			return err
		}
		return db.IncreaseVideoFavoriteCount(favorite.VideoID, 1)
	}
	if actionType == CancelFavorite {
		if err := db.RemoveFavorite(favorite.UserID, favorite.VideoID); err != nil {
			return errors.New("取消点赞失败")
		}
		return db.IncreaseVideoFavoriteCount(favorite.VideoID, -1)
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
