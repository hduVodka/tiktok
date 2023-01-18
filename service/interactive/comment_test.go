package interactive

import (
	"fmt"
	"testing"
	"tiktok/config"
	"tiktok/models"
)

func init() {
	config.Init()
	models.Init()
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

	status, err := CommentAction(c, 2)
	fmt.Println(status)
	if err != nil {
		t.Error(err)
	}
}

func TestCommentList(t *testing.T) {
	lis, err := CommentList(1)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(lis); i++ {
		fmt.Printf("%#v", *lis[i])
	}
}
