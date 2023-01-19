package models

import (
	"testing"
	"tiktok/config"
)

func init() {
	config.Init()
	Init()
}

func TestInsertUser(t *testing.T) {
	user := &User{
		Nickname: "1",
		Username: "22@qq.com",
		Password: "3sdfaerwt",
	}
	InsertUser(user)
}

func TestCheckUsername(t *testing.T) {
	user := &User{
		Nickname: "1",
		Username: "22@qq.com",
		Password: "3sdfaerwt",
	}
	CheckUsername(user)
}

func TestSearchUser(t *testing.T) {
	user := &User{
		Nickname: "1",
		Username: "22@qq.com",
		Password: "3sdfaerwt",
	}
	SearchUser(user)
}
