package video

import (
	"context"
	"fmt"
	"mime/multipart"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/log"
	"tiktok/models"
	"time"
)

func GetFeed(ctx context.Context, latestTime time.Time) ([]dto.Video, time.Time, error) {
	videos, err := db.GetFeedByTime(latestTime)
	if err != nil {
		return nil, time.Now(), err
	}

	var userId uint
	if ctx.Value("userId") != nil {
		userId = 0
	}
	userId, ok := ctx.Value("userId").(uint)
	if !ok {
		userId = 0
	}
	oldest := time.Now()
	res, err := modelVideos2dtoVideos(ctx, userId, videos)
	if err != nil {
		return nil, time.Now(), err
	}

	if len(videos) != 0 {
		oldest = videos[len(videos)-1].CreatedAt
	}

	return res, oldest, nil
}

func Publish(ctx context.Context, file multipart.File, ext string, title string) {
	// 大文件上传需要大量时间，不尽早返回客户端会超时
	go func() {
		filename, err := Upload(ctx, ext, file)
		if err != nil {
			log.Fatalln(err)
		}
		video := &models.Video{
			AuthorId: ctx.Value("userId").(uint),
			Title:    title,
			PlayUrl:  fmt.Sprintf("video/%s.mp4", filename),
			CoverUrl: fmt.Sprintf("cover/%s.jpg", filename),
		}

		// 这里应该用callback做但是现在开发环境没公网
		time.AfterFunc(time.Second*30, func() {
			for {
				if isExist(context.Background(), video.CoverUrl) {
					break
				}
				time.Sleep(time.Second * 10)
			}
			err = db.InsertVideo(video)
			if err != nil {
				log.Errorln(err)
			}
		})
	}()
}

func PublishList(ctx context.Context) ([]dto.Video, error) {
	var userId = ctx.Value("userId").(uint)
	videos, err := db.GetVideoListById(userId)
	if err != nil {
		return nil, err
	}
	return modelVideos2dtoVideos(ctx, userId, videos)
}

func modelVideos2dtoVideos(ctx context.Context, userId uint, videos []models.Video) ([]dto.Video, error) {
	res := make([]dto.Video, len(videos))
	for i, v := range videos {
		author := db.GetUser(v.AuthorId)
		isFav, err := db.IsFavorite(userId, v.ID)
		if err != nil {
			log.Errorf("fail to check favorite:%v", err)
		}
		res[i] = dto.Video{
			Id: v.ID,
			Author: dto.User{
				Id:            author.ID,
				Name:          author.Username,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       GenUrl(ctx, v.PlayUrl),
			CoverUrl:      GenUrl(ctx, v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFav,
			Title:         v.Title,
		}
	}
	return res, nil
}
