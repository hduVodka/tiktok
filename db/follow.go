package db

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiktok/log"
	"tiktok/models"
)

func IsFollow(ctx context.Context, userId uint, toUserId uint) bool {
	if err := db.Where("user_id = ? AND to_user_id = ?", userId, toUserId).First(&models.Follow{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Fatalln(err)
		}
		return false
	}
	return true
}

func InsertFollow(ctx context.Context, userId uint, toUserId uint) error {
	if err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&models.Follow{
		UserId:   userId,
		ToUserId: toUserId,
	}).Error; err != nil {
		log.Error("insert follow error:", err)
		return ErrDatabase
	}
	return nil
}

func DeleteFollow(ctx context.Context, userId uint, toUserId uint) error {
	if err := db.Where("user_id = ? AND to_user_id = ?", userId, toUserId).Delete(&models.Follow{}).Error; err != nil {
		log.Error("delete follow error:", err)
		return ErrDatabase
	}
	return nil
}

func GetFollowListByUserID(ctx context.Context, userId uint) ([]models.User, error) {
	var user []models.User
	var ids []uint
	if err := db.Model(&models.Follow{}).Select("to_user_id").Where("user_id= ? ", userId).Find(&ids).Error; err != nil {
		log.Error("get follow list error:", err)
		return nil, ErrDatabase
	}
	if err := db.Where("id IN ?", ids).Find(&user).Error; err != nil {
		log.Error("get follow list error:", err)
		return nil, ErrDatabase
	}
	return user, nil
}

func GetFanListByUserId(ctx context.Context, userId uint) ([]models.User, error) {
	var user []models.User
	var ids []uint
	if err := db.Model(&models.Follow{}).Select("user_id").Where("to_user_id = ?", userId).Find(&ids).Error; err != nil {
		log.Error("get fan list error", err)
		return nil, ErrDatabase
	}
	if err := db.Where("id IN ?", ids).Find(&user).Error; err != nil {
		log.Error("get fan list error:", err)
		return nil, ErrDatabase
	}
	return user, nil
}

func GetFriendListByUserId() ([]models.User, error) {
	//todo: get friend list
	return []models.User{}, nil
}
