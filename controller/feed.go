package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dto"
	"tiktok/service/video"
	"time"
)

type FeedResp struct {
	Resp
	VideoList []dto.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	var latestTime time.Time
	unix, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInvalidParams,
		})
		return
	}

	if unix != 0 {
		latestTime = time.Unix(unix, 0)
	} else {
		latestTime = time.Now()
	}

	list, nextTime, err := video.GetFeed(c, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("internal server error:%v", err),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResp{
		Resp: Resp{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: list,
		NextTime:  nextTime.Unix(),
	})
}
