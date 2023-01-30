package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"tiktok/config"
	"tiktok/db"
	"tiktok/dto"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
	"time"
)

const videoPathFormat = "public/video/%s/%s.mp4"
const coverPathFormat = "public/cover/%s/%s.jpg"

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

	if len(videos) != 0 {
		oldest = videos[len(videos)-1].CreatedAt
	}

	return res, oldest, nil
}

func Publish(ctx context.Context, file multipart.File, title string) error {
	t := time.Now().Format("2006-01-02")
	uu := uuid.New()
	videoPath := fmt.Sprintf(videoPathFormat, t, uu.String())
	for {
		if _, err := os.Stat(videoPath); errors.Is(err, os.ErrNotExist) {
			break
		}
		uu = uuid.New()
		videoPath = fmt.Sprintf(videoPathFormat, t, uu.String())
	}

	filename := uu.String()
	err := os.Mkdir("public/video/"+t, 0770)
	if err != nil && !os.IsExist(err) {
		log.Errorf("fail to mkdir:%v", err)
		return errors.New("fail to create file")
	}

	dst, err := os.Create(fmt.Sprintf(videoPathFormat, t, filename))
	if err != nil {
		log.Errorf("fail to create file:%v", err)
		return errors.New("fail to create file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("fail to save file:%v", err)
	}

	coverPath := fmt.Sprintf(coverPathFormat, t, filename)
	if err = utils.GetCoverOfVideo(videoPath, coverPath); err != nil {
		return fmt.Errorf("fail to get cover of video:%v", err)
	}

	video := &models.Video{
		AuthorId: ctx.Value("userId").(uint),
		Title:    title,
		PlayUrl:  fmt.Sprintf("%s/%s.mp4", t, filename),
		CoverUrl: fmt.Sprintf("%s/%s.jpg", t, filename),
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
	host := config.Conf.GetString("server.host")
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
			PlayUrl:       GenUrl(host, "static/video", v.PlayUrl),
			CoverUrl:      GenUrl(host, "static/cover", v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFav,
			Title:         v.Title,
		}
	}
	return res, nil
}

func GenUrl(host string, path ...string) string {
	res := fmt.Sprintf("http://%s", host)
	for _, p := range path {
		res = fmt.Sprintf("%s/%s", res, p)
	}
	return res
}
