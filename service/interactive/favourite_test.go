package interactive

import (
	"testing"
	"tiktok/config"
	"tiktok/models"
)

func init() {
	config.Init()
	models.Init()
}

func TestFavoriteAction(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	_, err := FavoriteAction(f, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestFavoriteList(t *testing.T) {
	f := &models.Favorite{
		UserID: 1,
	}
	vs, err := FavoriteList(f)
	if err != nil {
		t.Error(err)
	}
	t.Log(len(vs))
}
