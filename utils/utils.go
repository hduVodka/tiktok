package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"tiktok/config"
	"time"
)

func Init() {
	// load jwt secret
	jwtSecret = []byte(config.Conf.GetString("auth.jwt_key"))
	// rand
	rand.Seed(time.Now().UnixMicro())
	// init cos client
	InitCos()
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
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

func GenerateSalt() string {
	b := make([]byte, 8)
	// A rand.Int63() generates 63 random bits
	for i, cache := 7, src.Int63(); i >= 0; i-- {
		b[i] = letters[int(cache)%len(letters)]
		cache >>= 6
	}
	return string(b[:])
}

func Gravatar(email string) string {
	hash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(email))))
	return "https://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
}
