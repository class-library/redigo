package redigo

import (
	"time"

	"github.com/go-redsync/redsync"
)

func NewLock(key string, expiry, retryDelay time.Duration, tries int) *redsync.Mutex {
	return RedisLock.NewMutex(key, redsync.SetExpiry(expiry), redsync.SetRetryDelay(retryDelay), redsync.SetTries(tries))
}
