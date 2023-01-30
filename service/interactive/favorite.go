package interactive

import (
	"errors"
	"tiktok/db"
	"tiktok/models"
)

const AddFavorite = 1
const CancelFavorite = 2

func FavoriteAction(favorite *models.Favorite, actionType int) (int, error) {
	// 执行点赞或取消点赞操作
	if actionType == AddFavorite {
		if err := db.AddFavorite(favorite); err != nil {
			return -1, err
		}
		return 0, nil
	} else if actionType == CancelFavorite {
		if err := db.RemoveFavorite(favorite.UserID, favorite.VideoID); err != nil {
			return -1, errors.New("取消点赞失败")
		}
		return 0, nil
	} else {
		return -1, errors.New("请求参数错误")
	}
}

func FavoriteList(favorite *models.Favorite) ([]models.Video, error) {
	var videoList []models.Video
	var err error

	if videoList, err = db.GetFavoriteListByUserID(favorite.UserID); err != nil {
		return nil, err
	}

	return videoList, nil

}
