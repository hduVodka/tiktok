package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service/interactive"
)

func CommentAction(c *gin.Context) {
	var userId interface{}

	userId, _ = c.Get("userId")

	var actionType int
	var videoId int
	var content string
	var err error

	if actionType, err = strconv.Atoi(c.Query("action_type")); err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			"请求参数错误",
		})
		return
	}

	if videoId, err = strconv.Atoi(c.Query("video_id")); err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			"请求参数错误",
		})
		return
	}

	content = c.Query("content")

	comment := &models.Comment{
		VideoID: uint(videoId),
		UserID:  userId.(uint),
		Content: content,
	}
	if code, err := interactive.CommentAction(comment, actionType); err != nil {
		c.JSON(code, Resp{
			code,
			err.Error(),
		})
	} else {
		c.JSON(code, Resp{
			code,
			"操作成功",
		})
	}
}

func CommentList(c *gin.Context) {
	var videoId int
	var err error
	if videoId, err = strconv.Atoi(c.Query("video_id")); err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			"请求参数错误",
		})
	}

	if commentList, err := interactive.CommentList(uint(videoId)); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			500,
			"获取评论列表失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode":   200,
			"StatusMsg":    "获取评论列表成功",
			"Comment_list": commentList,
		})
	}
}
