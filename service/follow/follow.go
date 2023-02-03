package follow

import (
	"context"
	"errors"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
	"tiktok/utils"
)

func RelationAction(ctx context.Context, toUserId uint, actionType int) error {
	userId := ctx.Value("userId").(uint)
	exist := db.IsFollow(ctx, userId, toUserId)

	if exist && actionType == 1 || !exist && actionType == 2 {
		return nil
	}

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

func FriendList(ctx context.Context, userId uint) ([]dto.FriendUser, error) {
	friendList, err := db.GetFanListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	users := make([]dto.FriendUser, len(friendList))
	for i, v := range friendList {
		msg, err := db.LatestMessageBetween(ctx, userId, v.ID)
		if err != nil {
			return nil, err
		}
		msgStr := ""
		var msgType int64 = dto.RECEIVE
		if msg != nil {
			msgStr = msg.Content
		}
		if msg.UserId == userId {
			msgType = dto.SEND
		}
		users[i] = dto.FriendUser{
			User: dto.User{
				Id:            v.ID,
				Name:          v.Username,
				FollowCount:   v.FollowCount,
				FollowerCount: v.FollowerCount,
				IsFollow:      db.IsFollow(ctx, userId, v.ID),
				Avatar:        utils.Gravatar(v.Username),
			},
			Message: msgStr,
			MsgType: msgType,
		}
	}
	return users, nil
}
