package middlewave

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/utils"
)

type Resp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func Auth(c *gin.Context) {
	tokenString := c.Query("token")
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	if utils.CheckToken(userId, tokenString) {
		c.Set("userId", userId)
		c.Next()
	} else {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  "unauthorized",
		})
		c.Abort()
	}
}
