package interactive

import (
	"context"
	"fmt"
	"testing"
	"tiktok/config"
	"tiktok/db"
	"tiktok/models"
)

func init() {
	config.Init()
	db.Init()
}

func TestCommentAction(t *testing.T) {
	c := &models.Comment{
		UserID:  1,
		VideoID: 1,
		Content: "hello",
	}
	/*status, err := CommentAction(c, 1)
	fmt.Println(status)
	if err != nil {
	 t.Error(err)
	}*/

	err := CommentAction(c, 2)
	if err != nil {
		t.Error(err)
	}
}

func TestCommentList(t *testing.T) {
	var ctx context.Context
	lis, err := CommentList(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(lis); i++ {
		fmt.Printf("%#v", lis[i])
	}
}
