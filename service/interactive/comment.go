package interactive

import (
	"context"
	"errors"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/models"
)

// CommentAction 写/改/删评论操作
func CommentAction(ctx context.Context, comment *models.Comment, actionType int) error {
	// 评论
	if actionType == 1 {
		if err := db.InsertComment(comment); err != nil {
			return errors.New("评论失败")
		}
		return db.IncreaseVideoCommentCount(ctx, comment.VideoID, 1)
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
		dtoList[i] = dto.Comment{
			Id: v.ID,
			User: dto.User{
				Id:            v.User.ID,
				Name:          v.User.Username,
				FollowCount:   v.User.FollowCount,
				FollowerCount: v.User.FollowerCount,
				IsFollow:      db.IsFollow(ctx, userId, v.UserID),
			},
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		}
	}
	return dtoList
}
