package db

import (
	"testing"
	"tiktok/models"
)

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
