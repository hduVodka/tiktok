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
	VideoList []dto.Video `json:"videoList,omitempty"`
	NextTime  int64       `json:"nextTime,omitempty"`
}

func Feed(c *gin.Context) {
	var latestTime time.Time
	unix, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			StatusCode: http.StatusBadRequest,
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
		c.JSON(http.StatusInternalServerError, Resp{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  fmt.Sprintf("internal server error:%v", err),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResp{
		Resp: Resp{
			StatusCode: http.StatusOK,
			StatusMsg:  "ok",
		},
		VideoList: list,
		NextTime:  nextTime.Unix(),
	})
}
