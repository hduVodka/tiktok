package interactive

import (
	"context"
	"testing"
	"tiktok/config"
	"tiktok/db"
	"tiktok/models"
)

func init() {
	config.Init()
	db.Init()
}

func TestFavoriteAction(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	err := FavoriteAction(f, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestFavoriteList(t *testing.T) {
	vs, err := FavoriteList(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(len(vs))
}
