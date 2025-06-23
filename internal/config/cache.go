package config

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitCache(context context.Context) {
	ctx = context
	conf := GetCache()
	cache = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.Database,
		PoolSize: 50000,
	})
}

func SaveCache(key string, value interface{}) error {
	err := cache.Set(ctx, key, value, 60*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func ReadCache(key string) ([]byte, error) {
	val, err := cache.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func ScanCache(scanKey string) ([]string, error) {
	var keysResult []string
	var cursor uint64
	var keys []string
	var err error

	for {
		keysResult, cursor, err = cache.Scan(ctx, cursor, scanKey, 10).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, keysResult...)

		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

func MReadCache(keys ...string) ([][]byte, error) {
	if len(keys) == 0 {
		return [][]byte{}, nil
	}

	val, err := cache.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	byteVal := make([][]byte, len(val))
	for i, v := range val {
		byteVal[i] = []byte(v.(string))
	}

	return byteVal, nil
}

func DeleteCache(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	err := cache.Del(ctx, keys...).Err()
	if err != nil {
		return err
	}
	return nil
}

func CheckCacheHealth() error {
	_, err := cache.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func CountCache() (int64, error) {
	count, err := cache.DBSize(ctx).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
