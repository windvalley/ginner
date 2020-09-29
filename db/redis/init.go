package redis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"

	"ginner/config"
)

// Redis instance of redis pool
var Redis *redis.Pool

// Init redis pool initialization
func Init() {
	password := config.Conf().Redis.Password
	Redis = &redis.Pool{
		MaxIdle:     config.Conf().Redis.MaxIdle,
		MaxActive:   config.Conf().Redis.MaxActive,
		IdleTimeout: config.Conf().Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Conf().Redis.Address)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", config.Conf().Redis.DB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Set set key/value
func Set(key string, data interface{}, time int) error {
	conn := Redis.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists if the key exists
func Exists(key string) bool {
	conn := Redis.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get the value of a key
func Get(key string) ([]byte, error) {
	conn := Redis.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete key/value
func Delete(key string) (bool, error) {
	conn := Redis.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// BatchDelete delete some key/values by fuzzy matching
func BatchDelete(key string) error {
	conn := Redis.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
