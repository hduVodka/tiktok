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
}

func TestFavorite_InsertFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 1,
	}
	InsertFavorite(f)
}

func TestFavorite_DeleteFavorite(t *testing.T) {
	err := DeleteFavorite(1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestFavorite_GetFavoriteListByUserID(t *testing.T) {
	vs, err := GetFavoriteListByUserID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(len(vs))
}

func TestFavorite_IsFavorite(t *testing.T) {
	t.Log(IsFavorite(1, 1))
}

func TestCountFavoriteByID(t *testing.T) {
	count, err := CountFavoriteByID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
}
