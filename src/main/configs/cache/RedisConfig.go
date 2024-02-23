package main_configs_cache

import (
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"sync"
)

const _MSG_ATTEMPT_TO_CONNECT_TO_REDIS_CLIENT = "Attempt to connect to redis client %s"
const _MSG_SUCESSFULLY_CONNECTED_TO_REDIS_CLIENT = "Successfully connected to redis client %s"

var redisClientBean *redis.Client
var once sync.Once

func GetRedisClusterBean() *redis.Client {
	once.Do(func() {

		if redisClientBean == nil {
			redisClientBean = getRedisCluster()
		}

	})
	return redisClientBean
}

func getRedisCluster() *redis.Client {
	redisHost := main_configs_yml.GetYmlValueByName(main_configs_yml.RedisHosts)

	slog.Info(_MSG_ATTEMPT_TO_CONNECT_TO_REDIS_CLIENT, redisHost)

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping(context.TODO()).Err()
	if err != nil {
		panic(err)
	}

	slog.Info(_MSG_SUCESSFULLY_CONNECTED_TO_REDIS_CLIENT, redisHost)
	return client
}
