package middlewave

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/utils"
)

func Auth(c *gin.Context) {
	tokenString := c.Query("token")
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	if utils.CheckToken(userId, tokenString) {
		c.Set("userId", userId)
		c.Next()
	} else {
		c.Abort()
	}
}
