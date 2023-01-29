package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

var ffmpegPath string

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
