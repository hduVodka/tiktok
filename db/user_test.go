package db

import (
	"testing"
	"tiktok/config"
	"tiktok/models"
)

func init() {
	config.Init()
	Init()
}

func TestInsertUser(t *testing.T) {
	user := &models.User{
		Nickname: "1",
		Username: "22@qq.com",
		Password: "3sdfaerwt",
	}
	InsertNewUser(user)
}

func TestSearchUser(t *testing.T) {
	user := &models.User{
		Nickname: "1",
		Username: "22@qq.com",
		Password: "3sdfaerwt",
	}
	SearchUser(user)
}
