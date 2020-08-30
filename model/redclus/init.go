package redclus

import (
	"github.com/go-redis/redis"

	"use-gin/config"
)

var Redis *redis.ClusterClient

func Init() {
	Redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       config.Conf().RedisCluster.Nodes,
		Password:    config.Conf().RedisCluster.Secret,
		PoolSize:    config.Conf().RedisCluster.PoolSize,
		PoolTimeout: config.Conf().RedisCluster.PoolTimeout,
	})
}
