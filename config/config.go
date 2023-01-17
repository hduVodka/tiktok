package config

import (
	"bytes"
	_ "embed"
	vp "github.com/spf13/viper"
	"os"
	"tiktok/log"
)

var Conf *vp.Viper

//go:embed config_sample.yml
var defaultConfig []byte

func Init() {
	Conf = vp.New()
	Conf.SetConfigType("yaml")
	Conf.SetConfigName("config")
	Conf.AddConfigPath(".")
	if err := Conf.ReadInConfig(); err != nil {
		if _, ok := err.(vp.ConfigFileNotFoundError); ok {
			if err := os.WriteFile("./config.yml", defaultConfig, 0666); err != nil {
				log.Fatalf("fail to write config:%v", err)
			}
			_ = Conf.ReadConfig(bytes.NewReader(defaultConfig))
		} else {
			log.Fatalf("fail to read config:%v", err)
		}
	}
}
