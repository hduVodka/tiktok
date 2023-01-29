package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dto"
	"tiktok/service/follow"
)

type ListResp struct {
	Resp
	List []dto.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	if err := follow.RelationAction(c, uint(toUserId), actionType); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  "succeed",
	})
}

func FollowList(c *gin.Context) {
	followList, err := follow.FollowList(c)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get follow list",
		})
		return
	}

	c.JSON(http.StatusOK, ListResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		List: followList,
	})
}

func FollowerList(c *gin.Context) {
	fanList, err := follow.FollowerList(c)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get",
		})
		return
	}

	c.JSON(http.StatusOK, ListResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		List: fanList,
	})
}

func FriendList(c *gin.Context) {
	friendList, err := follow.FriendList(c)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get",
		})
		return
	}

	c.JSON(http.StatusOK, ListResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		List: friendList,
	})
}
