package redigo

import (
	"fmt"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

type RedisServer struct {
	Pool *redis.Pool
}

var Redis *RedisServer

var RedisLock *redsync.Redsync

func NewRedis(options ...Option) {
	data := (&Options{}).init()

	for _, option := range options {
		option(data)
	}

	Redis = new(RedisServer)
	Redis.Pool = &redis.Pool{
		MaxIdle:     data.MaxIdle,
		MaxActive:   data.MaxActive,
		IdleTimeout: data.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				data.NetWork,
				fmt.Sprintf("%s:%d", data.Host, data.Port),
				redis.DialReadTimeout(data.ReadTimeout),
				redis.DialWriteTimeout(data.WriteTimeout),
				redis.DialConnectTimeout(data.ConnectTimeout),
				redis.DialPassword(data.Password),
				redis.DialDatabase(data.Database),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	RedisLock = redsync.New([]redsync.Pool{Redis.Pool})
}
