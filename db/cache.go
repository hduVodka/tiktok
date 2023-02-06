package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"tiktok/log"
	"tiktok/models"
	"time"
)

func CacheInit() {
	// video cache
	// video的缓存和其他不同，视频id全量缓存，不需要过期
	// 用于feed接口
	var videos []models.Video
	if err := db.Order("created_at DESC").Find(&videos).Error; err != nil {
		log.Fatalln("init video cache fail", err)
	}

	var cache []redis.Z
	for _, v := range videos {
		cache = append(cache, redis.Z{
			Score:  float64(v.CreatedAt.UnixMilli()),
			Member: v.ID,
		})
	}
	if err := rdb.ZAdd(context.Background(), "video:feed", cache...).Err(); err != nil {
		log.Fatalln("init video cache fail", err)
	}
}

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
