package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/log"
	"tiktok/service/video"
)

func Publish(c *gin.Context) {
	fh, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		log.Error(err)
		return
	}

	title, existed := c.GetPostForm("title")
	if !existed {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		log.Error(err)
		return
	}

	// todo:检查文件类型

	file, err := fh.Open()
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		log.Error(err)
		return
	}
	defer file.Close()

	if err := video.Publish(c, file, title); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("internal server error:%v", err),
		})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  "ok",
	})
}

func PublishList(c *gin.Context) {

}
