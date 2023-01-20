package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"tiktok/config"
	"tiktok/log"
)

var ffmpegPath string

func Init() {
	// check ffmpeg if existed
	ffmpegPath = config.Conf.GetString("video.ffmpeg")
	if ffmpegPath == "" {
		ffmpegPath = "ffmpeg"
	}
	_, err := exec.Command(ffmpegPath, "-version").CombinedOutput()
	if err != nil {
		log.Fatalln("ffmpeg is not existed, please install ffmpeg first")
	}
}

func GetCoverOfVideo(videoPath string, coverPath string) error {
	err := os.MkdirAll(path.Dir(coverPath), 0770)
	if err != nil {
		return err
	}
	out, err := exec.Command(ffmpegPath, "-i", videoPath, "-y", "-f", "image2", "-ss", "1", "-frames:v", "1", coverPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("fail to get cover of video: %v: %s", err, string(out))
	}
	return nil
}
