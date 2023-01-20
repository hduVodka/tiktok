package utils

import (
	"testing"
)

func TestGetCoverOfVideo(t *testing.T) {
	ffmpegPath = "ffmpeg"
	err := GetCoverOfVideo("public/video/test.mp4", "public/cover/test.jpg")
	if err != nil {
		t.Error(err)
		return
	}
}
