package models

import (
	"testing"
	"tiktok/config"
)

func init() {
	config.Init()
	Init()
}

func TestFollow_InsertFollow(t *testing.T) {
	f := &Follow{
		UserId:     1,
		ToUserId:   300,
		ActionType: 1,
	}
	f.InsertFollow()
}

func TestFollow_DeleteFollow(t *testing.T) {
	f := &Follow{
		UserId:   1,
		ToUserId: 200,
	}
	f.DeleteFollow()
}

func TestFollow_GetFollowListByUserID(t *testing.T) {
	f := &Follow{
		UserId: 1,
	}
	list, err := f.GetFollowListByUserID()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestFollow_GetFanListByUserId(t *testing.T) {
	f := &Follow{
		UserId: 1,
	}
	list, err := f.GetFanListByUserId()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestFollow_IsFollow(t *testing.T) {
	f := &Follow{
		ToUserId: 300,
	}
	t.Log(f.IsFollow())
}

func TestFollow_GetFriendListByUserId(t *testing.T) {
	f := &Follow{
		UserId: 1,
	}
	list, err := f.GetFriendListByUserId()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}
