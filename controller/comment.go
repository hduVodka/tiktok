package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
	"tiktok/service/interactive"
)

type CommentListResp struct {
	Resp
	CommentList []dto.Comment `json:"comment_list"`
}

func CommentAction(c *gin.Context) {
	var content string
	user := new(models.User)
	userId := c.GetUint("userId")
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}
	if actionType == 1 {
		content = c.Query("comment_text")
	}

	comment := &models.Comment{
		VideoID: uint(videoId),
		UserID:  userId,
		Content: content,
	}
	if err := interactive.CommentAction(c, comment, actionType); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	if content != "" {
		user = db.GetUser(userId)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "string",
			"comment": dto.Comment{
				Id: comment.ID,
				User: dto.User{
					Id:            user.ID,
					Name:          user.Username,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      true,
				},
				Content:    comment.Content,
				CreateDate: strconv.FormatInt(comment.CreateTime, 10),
			},
		})
	} else {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 0,
			StatusMsg:  "ok",
		})
	}

}

func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			-1,
			ErrInvalidParams,
		})
	}

	if commentList, err := interactive.CommentList(c, uint(videoId)); err != nil {
		c.JSON(http.StatusOK, Resp{
			-1,
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
