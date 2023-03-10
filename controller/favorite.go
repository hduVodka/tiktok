package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dto"
	"tiktok/models"
	"tiktok/service/interactive"
)

type FavoriteListResp struct {
	Resp
	VideoList []dto.Video `json:"video_list"`
}

func FavoriteAction(c *gin.Context) {
	// 从上下文中获取用户id
	userId := c.GetUint("userId")

	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	videoID, err := strconv.ParseUint(c.Query("video_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	favorite := &models.Favorite{
		VideoID: uint(videoID),
		UserID:  userId,
	}

	if err := interactive.FavoriteAction(c, favorite, actionType); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  "ok",
	})

}

func FavoriteList(c *gin.Context) {
	userId := c.GetUint("userId")

	if videoList, err := interactive.FavoriteList(c, userId); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 500,
			StatusMsg:  "获取点赞列表失败",
		})
	} else {
		c.JSON(http.StatusOK, FavoriteListResp{
			Resp: Resp{
				StatusCode: 0,
				StatusMsg:  "获取点赞列表成功",
			},
			VideoList: videoList,
		})
	}
}
