package middlewave

import (
	"github.com/gin-gonic/gin"
	"tiktok/utils"
)

type Resp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func Auth(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		tokenString = c.PostForm("token")
	}
	if tokenString == "" {
		c.AbortWithStatusJSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  "token is empty",
		})
		return
	}

	userId, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  "invalid jwt token",
		})
		return
	}

	c.Set("userId", uint(userId))
}
