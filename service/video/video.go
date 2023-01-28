package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
	"time"
)

const videoDir = "public/video/%s/%s.mp4"
const coverDir = "public/cover/%s/%s.jpg"

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
	res, err := modelVideos2dtoVideos(userId, videos)
	if err != nil {
		return nil, time.Now(), err
	}

	oldest = videos[len(videos)-1].CreatedAt

	return res, oldest, nil
}

func Publish(ctx context.Context, file multipart.File, title string) error {
	var videoPath, coverPath string
	t := time.Now().Format("2006-01-02")
	for {
		uu := uuid.New()
		videoPath = fmt.Sprintf(videoDir, t, uu.String())
		if _, err := os.Stat(videoPath); errors.Is(err, os.ErrNotExist) {
			coverPath = fmt.Sprintf(coverDir, t, uu.String())
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

	if err = utils.GetCoverOfVideo(videoPath, coverPath); err != nil {
		return fmt.Errorf("fail to get cover of video:%v", err)
	}

	video := &models.Video{
		AuthorId: ctx.Value("userId").(uint),
		Title:    title,
		PlayUrl:  videoPath,
		CoverUrl: coverPath,
	}
	err = db.InsertVideo(video)
	if err != nil {
		return err
	}

	return nil
}

func PublishList(c context.Context) ([]dto.Video, error) {
	var userId = c.Value("userId").(uint)
	videos, err := db.GetVideoListById(userId)
	if err != nil {
		return nil, err
	}
	return modelVideos2dtoVideos(userId, videos)
}

func modelVideos2dtoVideos(userId uint, videos []models.Video) ([]dto.Video, error) {
	res := make([]dto.Video, len(videos))
	for i, v := range videos {
		author := db.GetUser(v.AuthorId)
		isFav, err := db.IsFavorite(userId, v.ID)
		if err != nil {
			log.Error(err)
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
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFav,
			Title:         v.Title,
		}
	}
	return res, nil
}
