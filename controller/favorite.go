package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service/interactive"
)

func FavoriteAction(c *gin.Context) {

	// 从上下文中获取用户id
	var userID interface{}
	var exist bool

	if userID, exist = c.Get("userId"); !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "未登录",
		})
		return
	}

	// 解析请求参数
	var actionType int
	var videoID int
	var err error

	if actionType, err = strconv.Atoi(c.Query("action_type")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	if videoID, err = strconv.Atoi(c.Query("video_id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	favorite := &models.Favorite{
		VideoID: uint(videoID),
		UserID:  userID.(uint),
	}

	if code, err := interactive.FavoriteAction(favorite, actionType); err != nil {
		c.JSON(code, gin.H{
			"status_code": code,
			"status_msg":  err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"status_code": code,
			"status_msg":  "操作成功",
		})
	}
}

func FavoriteList(c *gin.Context) {
	var userID interface{}
	var exist bool

	if userID, exist = c.Get("userId"); !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "未登录",
		})
		return
	}

	favorite := &models.Favorite{
		UserID: userID.(uint),
	}

	if videoList, err := interactive.FavoriteList(favorite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg":  "获取点赞列表失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "获取点赞列表成功",
			"video_list":  videoList,
		})
	}
}
