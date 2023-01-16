package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tiktok/helper"
	"tiktok/models"
)

func UserInfo(c *gin.Context) {

}

func Register(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := models.Init()
	// check username
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "连接数据库失败"})
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}

	// hash password 存入的是加密的密码
	hashedPassword := helper.Md5(user.Password)
	user.Password = string(hashedPassword)

	// Insert user into the database
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "插入用户数据失败"})
		return
	}

	// 调用helper中的GenerateToken方法生成token
	token, _ := helper.GenerateToken(user.Username, user.Password)

	// 返回数据
	c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "token": token})
}

func Login(c *gin.Context) {

}
