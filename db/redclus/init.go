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
//    e.g.
// fooStr, _ := redclus.Get("rediskey")
// if fooStr != "" {
//     if err := json.Unmarshal([]byte(fooStr), &foo); err != nil {
//         return err
//     }
//     return nil
// }
func Get(key string) (string, error) {
	return Redis.Get(key).Result()
}

// Set redis set
//    e.g.
// fooBytes, err := json.Marshal(foo)
// if err != nil {
//     return err
// }
// if _, err = redclus.Set("rediskey", fooBytes, 120); err != nil {
//     return err
// }
func Set(key string, value interface{}, expiration int) (string, error) {
	return Redis.Set(key, value, time.Duration(expiration)*time.Second).Result()
}

// Del redis del
func Del(key string) (int64, error) {
	return Redis.Del(key).Result()
}
