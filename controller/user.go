package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
	"tiktok/service/auth"
	"tiktok/utils"
)

type UserResp struct {
	Resp
	Token string `json:"token,omitempty"`
	Id    uint   `json:"user_id,omitempty"`
}

type UserInfoResp struct {
	StatusCode int64    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	User       dto.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	userId := c.Keys["userId"].(uint)
	if user := db.GetUser(userId); user != nil {
		c.JSON(200, UserInfoResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User: dto.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
			},
		})
	}
}

func Register(c *gin.Context) {

	user := new(models.User)
	user.Username = c.Query("username")
	user.Password = c.Query("password")
	user.Nickname = c.Query("nickname")

	// 用户名密码合法性检查
	if !auth.CheckLegal(user) {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrFormatError,
		})
		return
	}

	// 入库
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
		Id:    user.ID,
		Token: utils.GenerateToken(user.ID),
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
		Id:    user.ID,
		Token: utils.GenerateToken(user.ID),
	})
}
