package db

import (
	"os"
	"testing"
	"tiktok/config"
	"tiktok/models"
)

func init() {
	// 修改工作目录，解决配置文件读取问题
	os.Chdir("../")
	config.Init()
	models.Init()
	//db = db.Debug()
}

func TestFavorite_InsertFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	InsertFavorite(f)
}

func TestFavorite_DeleteFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	DeleteFavorite(f)
}

func TestFavorite_GetFavoriteListByUserID(t *testing.T) {
	f := &models.Favorite{
		UserID: 1,
	}
	vs, err := GetFavoriteListByUserID(f)
	if err != nil {
		t.Error(err)
	}
	t.Log(len(vs))
}

func TestFavorite_IsFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	t.Log(IsFavorite(f))
}

func TestCountFavoriteByID(t *testing.T) {
	count, err := CountFavoriteByID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
}
