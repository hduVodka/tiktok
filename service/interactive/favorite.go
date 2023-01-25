package interactive

import (
	"errors"
	"net/http"
	"tiktok/db"
	"tiktok/models"
)

func FavoriteAction(favorite *models.Favorite, actionType int) (int, error) {
	var exist bool
	// 执行点赞或取消点赞操作
	if actionType == 1 {
		exist,_ = db.IsFavorite(favorite.UserID, favorite.VideoID)
		if exist {
			return http.StatusOK, nil
		} else {
			if err := db.InsertFavorite(favorite); err != nil {
				return http.StatusInternalServerError, errors.New("点赞失败")
			}
			return http.StatusOK, nil
		}
	} else if actionType == 2 {
		exist,_ = db.IsFavorite(favorite.UserID, favorite.VideoID)
		if !exist {
			return http.StatusOK, nil
		}

		if err := db.DeleteFavorite(favorite.UserID, favorite.VideoID); err != nil {
			return http.StatusInternalServerError, errors.New("取消点赞失败")
		}

		return http.StatusOK, nil
	} else {
		return http.StatusBadRequest, errors.New("请求参数错误")
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
