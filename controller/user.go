package controller

import (
	"github.com/gin-gonic/gin"
	"sync"
	"tiktok/db"
	"tiktok/models"
	"tiktok/service/auth"
	"tiktok/utils"
)

type UserResp struct {
	Resp
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type UserInfoResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}

type User struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

func UserInfo(c *gin.Context) {
	userId := c.Keys["userId"].(uint)
	user := new(models.User)
	if utils.FindUserInfo(userId, user) {
		c.JSON(200, UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: User{
				ID:            user.ID,
				Name:          user.Username,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
				IsFollow:      user.IsFollow,
			},
		})
	}
}

func Register(c *gin.Context) {

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	user := new(models.User)

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	// 用户名密码合法性检查
	if !auth.CheckLegal(user) {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrFormatError,
		})
		return
	}

	// check username
	if !db.CheckUsername(user) {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrUserAlreadyExist,
		})
		return
	}

	// 加密密码,同时创建生成salt，并入库
	err := db.InsertNewUser(auth.Encrypt(user))
	if err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInsertFailed,
		})
		return
	}

	// 返回id和token
	c.JSON(200, UserResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "register success",
		},
		UserId: user.ID,
		Token:  utils.GenerateToken(user),
	})
}

func Login(c *gin.Context) {

	user := new(models.User)

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	// 查找用户,若存在则返回token
	if !db.SearchUser(user) {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrIncorrectPassword,
		})
		return
	}

	// 返回id和token
	c.JSON(200, UserResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "login success",
		},
		UserId: user.ID,
		Token:  utils.GenerateToken(user),
	})
}
