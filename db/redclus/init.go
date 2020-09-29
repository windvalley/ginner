package redclus

import (
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
