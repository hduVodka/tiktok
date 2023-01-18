package models

import (
	"testing"
	"tiktok/config"
)

func init() {
	config.Init()
	Init()
	//db = db.Debug()
}

func TestFavorite_InsertFavorite(t *testing.T) {
	f := &Favorite{
		UserID:  1,
		VideoID: 1,
	}
	f.InsertFavorite()
}

func TestFavorite_DeleteFavorite(t *testing.T) {
	f := &Favorite{
		UserID:  1,
		VideoID: 1,
	}
	f.DeleteFavorite()
}

func TestFavorite_GetFavoriteListByUserID(t *testing.T) {
	f := &Favorite{
		UserID: 1,
	}
	vs, err := f.GetFavoriteListByUserID()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(vs))
}

func TestFavorite_IsFavorite(t *testing.T) {
	f := &Favorite{
		UserID:  1,
		VideoID: 1,
	}
	t.Log(f.IsFavorite())
}

func TestCountFavoriteByID(t *testing.T) {
	count, err := CountFavoriteByID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
}
