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
	Init()
}

func TestFavorite_InsertFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 2,
	}
	err := AddFavorite(f)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestFavorite_DeleteFavorite(t *testing.T) {
	err := RemoveFavorite(1, 2)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFavorite_GetFavoriteListByUserID(t *testing.T) {
	vs, err := GetFavoriteListByUserID(1)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(len(vs))
}

func TestFavorite_IsFavorite(t *testing.T) {
	t.Log(IsFavorite(1, 2))
}

func TestTemp(t *testing.T) {
	deleteCache()
	TestFavorite_InsertFavorite(t)

	TestFavorite_GetFavoriteListByUserID(t)
}
