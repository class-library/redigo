package redigo

import (
	"github.com/go-redsync/redsync"
)

// NewLock 实例化Redis互斥锁
func NewLock(key string, options ...LockOption) *redsync.Mutex {
	data := (&LockOptions{}).init()

	for _, option := range options {
		option(data)
	}

	var redsyncOptions []redsync.Option

	if data.Tries > 0 {
		redsyncOptions = append(redsyncOptions, redsync.SetTries(data.Tries))
	}
	if data.Expiry > 0 {
		redsyncOptions = append(redsyncOptions, redsync.SetExpiry(data.Expiry))
	}
	if data.RetryDelay > 0 {
		redsyncOptions = append(redsyncOptions, redsync.SetRetryDelay(data.RetryDelay))
	}

	return RedisLock.NewMutex(key, redsyncOptions...)
}
