package video

import (
	"context"
	"fmt"
	"mime/multipart"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
	"time"
)

func GetFeed(ctx context.Context, latestTime time.Time) ([]dto.Video, time.Time, error) {
	videos, err := db.GetFeedByTime(ctx, latestTime)
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
	res, err := dto.FromVideoModels(ctx, userId, videos)
	if err != nil {
		return nil, time.Now(), err
	}

	if len(videos) != 0 {
		oldest = videos[len(videos)-1].CreatedAt
	}

	return res, oldest, nil
}

func Publish(ctx context.Context, fh *multipart.FileHeader, ext string, title string) {
	// 大文件上传需要大量时间，不尽早返回客户端会超时
	go func() {
		file, err := fh.Open()
		if err != nil {
			log.Errorln("fail to open uploaded file")
			return
		}
		defer file.Close()
		filename, err := utils.Upload(ctx, ext, file)
		if err != nil {
			log.Errorln(err)
		}
		video := &models.Video{
			AuthorId: ctx.Value("userId").(uint),
			Title:    title,
			PlayUrl:  fmt.Sprintf("video/%s.mp4", filename),
			CoverUrl: fmt.Sprintf("cover/%s.jpg", filename),
		}

		// 这里应该用callback做但是现在开发环境没公网
		for {
			time.Sleep(time.Second * 10)
			if utils.IsExist(context.Background(), video.CoverUrl) {
				break
			}
		}
		err = db.InsertVideo(video)
		if err != nil {
			log.Errorln(err)
		}
	}()
}

func PublishList(ctx context.Context) ([]dto.Video, error) {
	var userId = ctx.Value("userId").(uint)
	videos, err := db.GetVideoListById(userId)
	if err != nil {
		return nil, err
	}
	return dto.FromVideoModels(ctx, userId, videos)
}
