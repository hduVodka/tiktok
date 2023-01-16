package controller

import "github.com/gin-gonic/gin"

type FeedResp struct {
	Resp
	VideoList []Video `json:"videoList,omitempty"`
	NextTime  int64   `json:"nextTime,omitempty"`
}

func Feed(c *gin.Context) {

}
