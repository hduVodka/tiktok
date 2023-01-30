package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"tiktok/dto"
	"tiktok/log"
	"tiktok/service/video"
)

type PublishListResp struct {
	Resp
	VideoList []dto.Video `json:"videoList,omitempty"`
}

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

	if len(fh.Filename) == 0 {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		log.Error(err)
		return
	}

	ext := path.Ext(fh.Filename)

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

	video.Publish(c, file, ext, title)

	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  "ok",
	})
}

func PublishList(c *gin.Context) {
	videos, err := video.PublishList(c)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("internal server error:%v", err),
		})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, PublishListResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: videos,
	})
}
