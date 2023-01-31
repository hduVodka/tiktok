package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dto"
	"tiktok/service/message"
)

type MessageChatResp struct {
	Resp
	MessageList []dto.Message `json:"message_list"`
}

func MessageChat(c *gin.Context) {
	userId := c.GetUint("userId")

	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	messageList, err := message.GetList(c, userId, uint(toUserId))
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInternalServer,
		})
	}
	c.JSON(http.StatusOK, MessageChatResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		MessageList: messageList,
	})

}

func MessageAction(c *gin.Context) {
	userId := c.GetUint("userId")

	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	actionType, err := strconv.ParseUint(c.Query("action_type"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	content := c.Query("content")
	if content == "" {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	switch actionType {
	case 1:
		err := message.Send(c, userId, uint(toUserId), content)
		if err != nil {
			c.JSON(http.StatusOK, Resp{
				StatusCode: -1,
				StatusMsg:  ErrInternalServer,
			})
			return
		}
	default:
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	c.JSON(200, Resp{
		StatusCode: 0,
		StatusMsg:  "ok",
	})

}
