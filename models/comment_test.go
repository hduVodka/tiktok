package models

import (
	"fmt"
	"testing"
	"tiktok/config"
)

func init() {
	config.Init()
	Init()
}

func TestComment_InsertComment(t *testing.T) {
	c1 := &Comment{
		UserID:  3,
		VideoID: 1,
		Content: "评论3",
	}
	c1.InsertComment()
}

func TestComment_UpdateComment(t *testing.T) {
	c1 := &Comment{
		UserID:  1,
		VideoID: 1,
		Content: "评论7",
	}
	c1.UpdateComment()
}

func TestComment_DeleteComment(t *testing.T) {
	c := &Comment{
		UserID:  3,
		VideoID: 1,
		Content: "评论3",
	}
	c.DeleteComment()
}

func TestComment_GetCommentListByVideoId(t *testing.T) {
	lis, err := GetCommentListByVideoId(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(lis); i++ {
		fmt.Printf("%#v", lis[i])
	}
}

func TestCountCommentById(t *testing.T) {
	count, err := CountCommentById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", count)
}
