package redclus

import (
	"time"

	"github.com/go-redis/redis"

	"ginner/config"
)

// Redis client instance of redis cluster
var Redis *redis.ClusterClient

// Init redis cluster initialization
func Init() {
	Redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       config.Conf().RedisCluster.Nodes,
		Password:    config.Conf().RedisCluster.Secret,
		PoolSize:    config.Conf().RedisCluster.PoolSize,
		PoolTimeout: config.Conf().RedisCluster.PoolTimeout,
	})
}

// Get redis get
func Get(key string) (string, error) {
	return Redis.Get(key).Result()
}

// Set redis set
func Set(key string, value interface{}, expiration int) (string, error) {
	return Redis.Set(key, value, time.Duration(expiration)).Result()
}

// Del redis del
func Del(key string) (int64, error) {
	return Redis.Del(key).Result()
}
