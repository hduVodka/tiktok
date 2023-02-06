package db

import (
	"context"
	"testing"
	"tiktok/config"
	"time"
)

func TestGetFeedByTime(t *testing.T) {
	config.Init()
	Init()
	feed, err := GetFeedByTime(context.Background(), time.Now())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(feed)
}
