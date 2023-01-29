package follow

import (
	"context"
	"errors"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
)

func RelationAction(ctx context.Context, toUserId uint, actionType int) error {
	userId := ctx.Value("userId").(uint)
	exist := db.IsFollow(ctx, userId, toUserId)
	if actionType == 1 && !exist {
		return db.InsertFollow(ctx, userId, toUserId)
	}

	if actionType == 2 && exist {
		return db.DeleteFollow(ctx, userId, toUserId)
	}

	return errors.New("invalid params")
}

func FollowList(ctx context.Context) ([]dto.User, error) {
	var modelList []models.User
	var err error
	if modelList, err = db.GetFollowListByUserID(ctx, ctx.Value("userId").(uint)); err != nil {
		return nil, err
	}
	users := make([]dto.User, len(modelList))
	for i, v := range modelList {
		users[i] = dto.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      true,
		}
	}
	return users, nil
}

func FollowerList(ctx context.Context) ([]dto.User, error) {
	userId := ctx.Value("userId").(uint)
	modelList, err := db.GetFanListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	users := make([]dto.User, len(modelList))
	for i, v := range modelList {
		users[i] = dto.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      db.IsFollow(ctx, userId, v.ID),
		}
	}
	return users, nil
}

func FriendList(ctx context.Context) ([]dto.User, error) {
	userId := ctx.Value("userId").(uint)
	friendList, err := db.GetFriendListByUserId()
	if err != nil {
		return nil, err
	}

	users := make([]dto.User, len(friendList))
	for i, v := range friendList {
		users[i] = dto.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      db.IsFollow(ctx, userId, v.ID),
		}
	}
	return users, nil
}
