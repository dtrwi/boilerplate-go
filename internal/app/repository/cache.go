package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type ICache interface {
	Get(key string, res interface{}) (found bool, err error)
	Set(key string, res interface{}, expiration time.Duration) (err error)
}

func NewCache(opt Option) ICache {
	return &cache{
		Option: opt,
	}
}

type cache struct {
	Option
}

var ctx = context.Background()

func (r *cache) Get(key string, res interface{}) (found bool, err error) {
	value, err := r.Cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return
	}

	if err != nil {
		logrus.Warnf("error get redis |%s - %+v", key, err)
		return
	}

	if res != nil {
		if err = json.Unmarshal([]byte(value), res); err != nil {
			return
		}
		return true, nil
	}

	return
}

func (r *cache) Set(key string, res interface{}, expiration time.Duration) (err error) {
	value, err := json.Marshal(res)
	if err != nil {
		logrus.Warnf("error set redis |%s - %+v", key, err)
		return
	}

	err = r.Cache.Set(ctx, key, value, expiration).Err()

	return
}
