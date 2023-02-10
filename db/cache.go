package db

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
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

// the following 2 functions is used to store and get an object that will not be changed in redis
// objects will be serialized and deserialized with gob

func setStaticCache(ctx context.Context, key string, val any, expireTime time.Duration) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(val); err != nil {
		log.Error("encode fail: ", err)
		return
	}
	if err := rdb.Set(ctx, key, buf.Bytes(), expireTime).Err(); err != nil {
		log.Error("set cache fail: ", err)
	}
}

// val should be a pointer
func getStaticCache(ctx context.Context, key string, val any) {
	buf, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		log.Error("get cache fail: ", err)
	}
	if err := gob.NewDecoder(bytes.NewReader(buf)).Decode(val); err != nil {
		log.Error("decode fail: ", err)
	}
}

func groupQuery[T any, P uint | string](ctx context.Context, f func(ctx context.Context, id P) *T, ids []P) []T {
	res := make([]T, len(ids))
	var wg sync.WaitGroup
	wg.Add(len(ids))
	for i, id := range ids {
		go func(i int, id P) {
			res[i] = *f(ctx, id)
			wg.Done()
		}(i, id)
	}
	wg.Wait()
	return res
}
