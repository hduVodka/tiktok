package interactive

import (
	"errors"
	"net/http"
	"tiktok/models"
)

// CommentAction 写/改/删评论操作
func CommentAction(comment *models.Comment, actionType int) (int, error) {
	var exist bool
	// 评论
	if actionType == 1 {
		exist = comment.IsInSQL()
		if err := comment.InsertComment(); err != nil {
			return http.StatusInternalServerError, errors.New("评论失败")
		}
		return http.StatusOK, nil
	} else if actionType == 2 {
		exist = comment.IsInSQL()
		if !exist {
			return http.StatusOK, nil
		}
		if err := comment.DeleteComment(); err != nil {
			return http.StatusInternalServerError, errors.New("取消评论失败")
		}
		return http.StatusOK, nil
	} else {
		return http.StatusBadRequest, errors.New("请求参数错误")
	}
}

// CommentList 返回评论列表
func CommentList(videoId uint) ([]*models.Comment, error) {
	var commentList []*models.Comment
	var err error
	if commentList, err = models.GetCommentListByVideoId(videoId); err != nil {
		return nil, err
	}
	return commentList, nil
}
