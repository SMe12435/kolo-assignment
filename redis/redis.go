package redis

import (
	"sync"

	"kolo-assignment/config"

	"github.com/go-redis/redis"
)

var clusterClient *redis.ClusterClient
var once sync.Once

func NewRedisClient(clusterAdd []string) *redis.ClusterClient {
	redisClusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: clusterAdd,
	})
	return redisClusterClient
}

func Init() {
	once.Do(func() {
		clusterAdd := []string{config.Get().RedisAddr}
		clusterClient = NewRedisClient(clusterAdd)
		_, err := clusterClient.Ping().Result()
		if err != nil {
			panic(err.Error())
		}
	})
}

func GetClient() *redis.ClusterClient {
	return clusterClient
}
