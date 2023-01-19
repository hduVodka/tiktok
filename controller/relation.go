package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	strconv "strconv"
	"tiktok/models"
	follow2 "tiktok/service/follow"
)

type ListResp struct {
	Resp
	List []models.User
}

func RelationAction(c *gin.Context) {
	var userId interface{}

	var actionType int
	var toUserId uint
	var err error

	if actionType, err = strconv.Atoi(c.Query("action_type")); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	toUserId = c.GetUint("to_user_id")

	follow := &models.Follow{
		UserId:   userId.(uint),
		ToUserId: toUserId,
	}

	if err := follow2.RelationAction(follow, actionType); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 0,
			StatusMsg:  "succeed",
		})
	}

}

func FollowList(c *gin.Context) {
	var userId interface{}

	follow := &models.Follow{
		UserId: userId.(uint),
	}

	if followList, err := follow2.FollowList(follow); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get",
		})
	} else {
		c.JSON(http.StatusOK, ListResp{
			Resp: Resp{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			List: followList,
		})
	}
}

func FollowerList(c *gin.Context) {
	var userId interface{}

	follow := &models.Follow{
		UserId: userId.(uint),
	}

	if fanList, err := follow2.FollowerList(follow); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get",
		})
	} else {
		c.JSON(http.StatusOK, ListResp{
			Resp: Resp{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			List: fanList,
		})
	}
}

func FriendList(c *gin.Context) {
	var userId interface{}

	follow := &models.Follow{
		UserId: userId.(uint),
	}

	if friendList, err := follow2.FriendList(follow); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  "failed to get",
		})
	} else {
		c.JSON(http.StatusOK, ListResp{
			Resp: Resp{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			List: friendList,
		})
	}
}
