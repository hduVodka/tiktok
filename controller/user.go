package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/db"
	"tiktok/models"
	"tiktok/service/auth"
	"tiktok/utils"
)

func UserInfo(c *gin.Context) {

}

func Register(c *gin.Context) {

	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check username
	if ok, err := db.CheckUsername(user); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	auth.Encrypt(user)

	// Insert user into the database
	err := db.InsertNewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用utils中的GenerateToken方法生成token
	token, _ := utils.GenerateToken(user)

	// 返回数据
	c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "token": token})
}

func Login(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用户,若存在则返回token
	if ok, err := db.SearchUser(user); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		token, _ := utils.GenerateToken(user) // 生成token
		c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "token": token})
	}
}
