package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os/exec"
	"tiktok/config"
	"tiktok/log"
)

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
	// load jwt secret
	jwtSecret = []byte(config.Conf.GetString("auth.jwt_key"))
}

// sha256加密
func SHA256(password, salt string) string {
	code := []byte(salt)
	saltedPassword := append([]byte(password), ':')
	saltedPassword = append(saltedPassword, code...)

	hash := sha256.Sum256(saltedPassword)
	return hex.EncodeToString(hash[:])
}

// 生成随机salt
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func GenerateSalt() string {
	b := make([]byte, 8)
	for i := 0; i < 8; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
