package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dto"
	"tiktok/models"
	"tiktok/service/interactive"
)

type CommentListResp struct {
	Resp
	CommentList []dto.Comment `json:"comment_list"`
}

func CommentAction(c *gin.Context) {
	userId := c.GetUint("userId")

	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			ErrInvalidParams,
		})
		return
	}
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			ErrInvalidParams,
		})
		return
	}

	content := c.Query("content")

	comment := &models.Comment{
		VideoID: uint(videoId),
		UserID:  userId,
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
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			400,
			ErrInvalidParams,
		})
	}

	if commentList, err := interactive.CommentList(c, uint(videoId)); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			500,
			"获取评论列表失败",
		})
	} else {
		c.JSON(http.StatusOK, CommentListResp{
			Resp: Resp{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			CommentList: commentList,
		})
	}
}
