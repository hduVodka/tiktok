package models

import (
	"gorm.io/gorm"
	"tiktok/log"
)

type Follow struct {
	gorm.Model
	UserId     uint `json:"user_id"`
	ToUserId   uint `json:"toUserId"`
	FanId      uint `json:"fanId"`
	FriendId   uint `json:"friendId"`
	ActionType int  `json:"actionType"`
}

func (f *Follow) IsFollow() bool {
	var count int64
	if err := db.Model(&Follow{}).Where("to_user_id=?", f.ToUserId).Count(&count).Error; err != nil {
		log.Error("count follow error: ", err)
		return false
	}
	return count > 0
}

func (f *Follow) InsertFollow() error {
	if err := db.Create(f).Error; err != nil {
		log.Error("insert follow error:", err)
		return err
	}
	return nil
}

func (f *Follow) DeleteFollow() error {
	if err := db.Delete(f, "user_id=? AND to_user_id=?", f.UserId, f.ToUserId).Error; err != nil {
		log.Error("delete follow error:", err)
		return err
	}
	return nil
}

func (f *Follow) GetFollowListByUserID() ([]User, error) {
	var follow []Follow
	var user []User
	var ids []uint
	if err := db.Where("user_id=? AND action_type=?", f.UserId, 1).Find(&follow).Error; err != nil {
		log.Error("get follow list error:", err)
		return nil, err
	}
	for _, f := range follow {
		ids = append(ids, f.ToUserId)
	}
	if err := db.Where("user_id in ?", ids).Find(&user).Error; err != nil {
		log.Error("get follow list error:", err)
		return nil, err
	}
	return user, nil
}

func (f *Follow) GetFanListByUserId() ([]User, error) {
	var fan []Follow
	var user []User
	var ids []uint
	if err := db.Where("to_user_id=? AND action_type=?", f.ToUserId, 1).Find(&fan).Error; err != nil {
		log.Error("get fan list error", err)
		return nil, err
	}
	for _, f := range fan {
		ids = append(ids, f.UserId)
	}
	if err := db.Where("user_id in (?)", ids).Find(&user).Error; err != nil {
		log.Error("get fan list error:", err)
		return nil, err
	}
	return user, nil
}

func (f *Follow) GetFriendListByUserId() ([]User, error) {
	var friend []Follow
	var user []User
	var ids []uint
	if err := db.Where("user_id=?", f.UserId).Find(&friend).Error; err != nil {
		log.Error("get friend list error", err)
		return nil, err
	}
	for _, f := range friend {
		ids = append(ids, f.FriendId)
	}
	if err := db.Where("id in (?)", ids).Find(&user).Error; err != nil {
		log.Error("get fan list error:", err)
		return nil, err
	}
	return user, nil
}
