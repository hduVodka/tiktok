package db

import (
	"context"
	"testing"
	"tiktok/models"
	"time"
)

func TestGetFeedByTime(t *testing.T) {
	feed, err := GetFeedByTime(context.Background(), time.Now())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(feed)
}

func TestInsertVideo(t *testing.T) {
	video := &models.Video{
		AuthorId: 1,
		PlayUrl:  "222",
		CoverUrl: "333",
	}
	err := InsertVideo(context.Background(), video)
	if err != nil {
		t.Error(err)
	}
}
