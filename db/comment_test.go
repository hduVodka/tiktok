package db

import (
	"fmt"
	"testing"
	"tiktok/models"
)

func TestComment_InsertComment(t *testing.T) {
	c := &models.Comment{
		UserID:  1,
		VideoID: 5,
		Content: "this is a test",
	}
	InsertComment(c)
}

func TestComment_UpdateComment(t *testing.T) {
	c := &models.Comment{
		UserID:  1,
		VideoID: 5,
		Content: "update message",
	}
	err := UpdateComment(c)
	if err != nil {
		t.Error(err)
	}
}

func TestComment_DeleteComment(t *testing.T) {
	c := &models.Comment{
		UserID:  3,
		VideoID: 3,
	}
	err := DeleteComment(c)
	if err != nil {
		t.Error(err)
	}
}

func TestComment_GetCommentListByVideoId(t *testing.T) {
	var videoId uint = 8
	comments, err := GetCommentListByVideoId(videoId)
	if err != nil {
		t.Error(err)
	}
	for _, v := range comments {
		fmt.Println(v.User)
	}
}

func TestComment_CountCommentById(t *testing.T) {
	var videoId uint = 5
	count, err := CountCommentById(videoId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)
}
