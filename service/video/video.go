package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"tiktok/dto"
	"tiktok/models"
	"time"
)

func GetFeed(ctx context.Context, latestTime time.Time) ([]dto.Video, time.Time, error) {
	videos, err := models.GetFeedByTime(latestTime)
	if err != nil {
		return nil, time.Now(), err
	}

	oldest := time.Now()
	res := make([]dto.Video, len(videos))
	for i, v := range videos {
		// todo cache favorite and comment count
		res[i] = dto.Video{
			Id:            v.ID,
			Author:        dto.User{},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    0,
			Title:         v.Title,
		}
		if i == len(videos)-1 {
			oldest = v.CreatedAt
		}
	}
	return nil, oldest, nil
}

func Publish(ctx context.Context, file multipart.File, title string) error {
	var videoPath, coverPath string
	t := time.Now().Format("2006-01-02")
	for {
		uu := uuid.New()
		videoPath = fmt.Sprintf("public/video/%s/%s.mp4", t, uu.String())
		if _, err := os.Stat(videoPath); errors.Is(err, os.ErrNotExist) {
			coverPath = fmt.Sprintf("public/video/%s/%s.jpg", t, uu.String())
			break
		}
	}

	dst, err := os.Create(videoPath)
	if err != nil {
		return fmt.Errorf("fail to create file:%v", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("fail to save file:%v", err)
	}

	// todo:获取视频封面并保存

	video := &models.Video{
		AuthorId: ctx.Value("userId").(uint),
		Title:    title,
		PlayUrl:  videoPath,
		CoverUrl: coverPath,
	}
	err = models.InsertVideo(video)
	if err != nil {
		return err
	}

	return nil
}
