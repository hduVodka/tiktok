package db

import (
	"context"
	"tiktok/log"
	"time"
)

// UpdateCache 延时双删
func UpdateCache(ctx context.Context, key string, dbFunc func() error) error {
	if err := rdb.Del(ctx, key).Err(); err != nil {
		return err
	}
	if err := dbFunc(); err != nil {
		return err
	}
	time.AfterFunc(time.Second, func() {
		if err := rdb.Del(ctx, key).Err(); err != nil {
			log.Error("delete cache error: ", err)
		}
	})
	return nil
}
