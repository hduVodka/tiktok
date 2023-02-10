package interactive

import (
	"context"
	"errors"
	"strconv"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
)

// CommentAction 写/改/删评论操作
func CommentAction(ctx context.Context, comment *models.Comment, actionType int) error {
	// 评论
	if actionType == 1 {
		var count int64
		var err error
		if count, err = db.InsertComment(comment); err != nil {
			return errors.New("评论失败")
		}
		if count == 0 {
			return db.IncreaseVideoCommentCount(ctx, comment.VideoID, 1)
		}
		return nil
	} else if actionType == 2 {
		if err := db.DeleteComment(comment); err != nil {
			return errors.New("取消评论失败")
		}
		return db.IncreaseVideoCommentCount(ctx, comment.VideoID, -1)
	} else {
		return errors.New("请求参数错误")
	}
}

// CommentList 返回评论列表
func CommentList(ctx context.Context, videoId uint) ([]dto.Comment, error) {
	commentList, err := db.GetCommentListByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	return CommentModels2dto(ctx, commentList), nil
}

func CommentModels2dto(ctx context.Context, models []models.Comment) []dto.Comment {
	userId := ctx.Value("userId").(uint)
	dtoList := make([]dto.Comment, len(models))
	for i, v := range models {
		user := db.GetUserMessage(v)
		dtoList[i] = dto.Comment{
			Id: v.ID,
			User: dto.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      db.IsFollow(ctx, userId, v.UserID),
			},
			Content:    v.Content,
			CreateDate: strconv.FormatInt(v.CreateTime, 10),
		}
	}
	return dtoList
}
