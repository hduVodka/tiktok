package follow

import (
	"errors"
	"tiktok/models"
)

func RelationAction(follow *models.Follow, actionType int) error {
	var exist bool
	if actionType == 1 {
		exist = follow.IsFollow()
		if exist {
			return nil
		} else {
			if err := follow.InsertFollow(); err != nil {
				return errors.New("failed")
			}
			return nil
		}
	} else if actionType == 2 {
		exist = follow.IsFollow()
		if !exist {
			return nil
		}
		if err := follow.DeleteFollow(); err != nil {
			return errors.New("failed")
		}
		return nil
	} else {
		return errors.New("invalid params")
	}
}

func FollowList(follow *models.Follow) ([]models.User, error) {
	var followList []models.User
	var err error
	if followList, err = follow.GetFollowListByUserID(); err != nil {
		return nil, err
	}
	return followList, nil
}

func FollowerList(follow *models.Follow) ([]models.User, error) {
	var fanList []models.User
	var err error
	if fanList, err = follow.GetFanListByUserId(); err != nil {
		return nil, err
	}
	return fanList, nil
}

func FriendList(follow *models.Follow) ([]models.User, error) {
	var friendList []models.User
	var err error
	if friendList, err = follow.GetFriendListByUserId(); err != nil {
		return nil, err
	}
	return friendList, nil
}
