package main_configs_distributedlock

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis"
	"sync"
)

var RedisCluster *redsync.Redsync = nil
var once sync.Once

func GetLockClientBean() *redsync.Redsync {
	once.Do(func() {

		if RedisCluster == nil {
			RedisCluster = getLockClient()
		}

	})
	return RedisCluster
}

func getLockClient() *redsync.Redsync {
	redisClient := main_configs_cache.GetRedisClusterBean()
	pool := goredis.NewPool(redisClient)
	rs := redsync.New(pool)
	return rs
}
