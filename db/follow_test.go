package db

import (
	"context"
	"testing"
	"tiktok/config"
)

func init() {
	config.Init()
	Init()
}

func TestInsertFollow(t *testing.T) {
	err := InsertFollow(context.Background(), 1, 200)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteFollow(t *testing.T) {
	err := DeleteFollow(context.Background(), 1, 200)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetFollowListByUserID(t *testing.T) {
	list, err := GetFollowListByUserID(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestGetFanListByUserId(t *testing.T) {
	list, err := GetFanListByUserId(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestIsFollow(t *testing.T) {
	t.Log(IsFollow(context.Background(), 1, 200))
}

func TestGetFriendListByUserId(t *testing.T) {
	list, err := GetFriendListByUserId()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}
