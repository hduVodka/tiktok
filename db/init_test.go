package db

import (
	"os"
	"tiktok/config"
)

func init() {
	// 修改工作目录，解决配置文件读取问题
	os.Chdir("../")
	config.Init()
	Init()
}
