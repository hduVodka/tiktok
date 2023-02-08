package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"tiktok/log"
	"tiktok/models"
	"tiktok/utils"
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

// updateCache 延时双删
func updateCache(ctx context.Context, key string, dbFunc func() error) error {
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

func setHashCache[T any](ctx context.Context, key string, val T, expire time.Duration) {
	if _, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, key, val)
		pipe.Expire(ctx, fmt.Sprintf(key, val), expireTime)
		return nil
	}); err != nil {
		log.Errorf("set cache fail: %v", err)
	}
}

func getHashCache[T any](ctx context.Context, key string) *T {
	mp, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		log.Errorf("get cache fail:%v", err)
		return nil
	}
	return utils.Scan[T](mp)
}
