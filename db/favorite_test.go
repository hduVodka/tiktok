package db

import (
	"context"
	"testing"
	"tiktok/models"
)

func TestFavorite_InsertFavorite(t *testing.T) {
	f := &models.Favorite{
		UserID:  1,
		VideoID: 2,
	}
	err := AddFavorite(context.Background(), f)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestFavorite_DeleteFavorite(t *testing.T) {
	err := RemoveFavorite(context.Background(), 1, 2)
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
