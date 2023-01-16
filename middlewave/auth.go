package middlewave

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// todo:鉴权中间件
	// 鉴权完成后将userid写入上下文
	// 未通过则打回
	c.Set("userId", 1)
}
