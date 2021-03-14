package util

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"ginner/db/redclus"
	"ginner/logger"
)

// RedisMutex redis distributed lock object
type RedisMutex struct {
	Redis *redis.ClusterClient
	// redis key
	MutexName string
	// redis value
	ValueUUID string
	// key expiration seconds
	Expiration time.Duration
	// task timeout seconds
	Timeout time.Duration
	// delete the redis key while get an unlock signal
	UnlockC chan int
	Err     error
}

// NewRedisMutex get a new redis mutex with given name and other parameters
func NewRedisMutex(mutexName string, expiration,
	timeout time.Duration) *RedisMutex {

	uuid, err := GenerateUUID()
	c := make(chan int)
	return &RedisMutex{
		Redis:      redclus.Redis,
		MutexName:  mutexName,
		ValueUUID:  uuid,
		Expiration: expiration,
		Timeout:    timeout,
		UnlockC:    c,
		Err:        err,
	}
}

// Lock locker lock
func (m *RedisMutex) Lock() error {
	locker := m.MutexName + ":" + m.ValueUUID
	ret := m.Redis.SetNX(m.MutexName, m.ValueUUID, m.Expiration*time.Second)
	if !ret.Val() {
		return fmt.Errorf("%s locked", m.MutexName)
	}

	timeoutC := make(chan bool)
	go func() {
		time.Sleep(m.Timeout * time.Second)
		timeoutC <- true
	}()

	go func() {
		for {
			select {
			case <-time.After((m.Expiration - 10) * time.Second):
				ret := m.Redis.SetXX(m.MutexName, m.ValueUUID, m.Expiration*time.Second)
				if !ret.Val() {
					logger.Log.Warnf("locker %s renewal failed: %s", locker, ret.Err())
				}
			case <-m.UnlockC:
				logger.Log.Debugf("locker %s unlocked, so stop renewval", locker)
				return
			case <-timeoutC:
				logger.Log.Warnf("locker %s timeout, so stop renewval", locker)
				return
			}
		}
	}()

	return nil
}

// Unlock locker unlock
func (m *RedisMutex) Unlock() error {
	m.UnlockC <- 1
	locker := m.MutexName + ":" + m.ValueUUID
	value, err := redclus.Get(m.MutexName)
	if err != nil || value == "" {
		return fmt.Errorf("%s unlock error: %v", locker, err)
	}
	if value != m.ValueUUID {
		return fmt.Errorf("locker '%s' is not owned by the current task", locker)
	}

	if _, err = redclus.Del(m.MutexName); err != nil {
		return fmt.Errorf("%s unlock error: %v", locker, err)
	}

	return nil
}
