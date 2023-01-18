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

func UserInfo(c *gin.Context) {

}

func Register(c *gin.Context) {

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	user := new(models.User)

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	// 用户名密码合法性检查
	if !auth.CheckLegal(user) {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrFormatError,
		})
		return
	}

	// check username
	if !db.CheckUsername(user) {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrUserAlreadyExist,
		})
		return
	}

	// 加密密码,同时创建生成salt，并入库
	err := db.InsertNewUser(auth.Encrypt(user))
	if err != nil {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrInsertFailed,
		})
		return
	}

	// 返回id和token
	c.JSON(200, UserResp{
		Resp: Resp{
			StatusCode: 200,
			StatusMsg:  "register success",
		},
		UserId: user.ID,
		Token:  utils.GenerateToken(user),
	})
}

func Login(c *gin.Context) {

	user := new(models.User)

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	// 查找用户,若存在则返回token
	if !db.SearchUser(user) {
		c.JSON(400, Resp{
			StatusCode: 400,
			StatusMsg:  ErrIncorrectPassword,
		})
		return
	}

	// 返回id和token
	c.JSON(200, UserResp{
		Resp: Resp{
			StatusCode: 200,
			StatusMsg:  "login success",
		},
		UserId: user.ID,
		Token:  utils.GenerateToken(user),
	})
}
