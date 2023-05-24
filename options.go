package redigo

import "time"

// Options REDIS配置数据
type Options struct {
	Host           string
	Port           int
	Password       string
	Database       int
	NetWork        string
	MaxIdle        int
	MaxActive      int
	IdleTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

type LockOptions struct {
	Expiry     time.Duration
	RetryDelay time.Duration
	Tries      int
}

type LockOption func(options *LockOptions)

type Option func(*Options)

func WithExpiry(expiry time.Duration) LockOption {
	return func(options *LockOptions) {
		options.Expiry = expiry
	}
}

func WithRetryDelay(retryDelay time.Duration) LockOption {
	return func(options *LockOptions) {
		options.RetryDelay = retryDelay
	}
}

func WithTries(tries int) LockOption {
	return func(options *LockOptions) {
		options.Tries = tries
	}
}

func WithHost(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func WithPort(port int) Option {
	return func(options *Options) {
		options.Port = port
	}
}

func WithPassword(password string) Option {
	return func(options *Options) {
		options.Password = password
	}
}

func WithDatabase(database int) Option {
	return func(options *Options) {
		options.Database = database
	}
}

func WithNetWork(netWork string) Option {
	return func(options *Options) {
		options.NetWork = netWork
	}
}

func WithMaxIdle(maxIdle int) Option {
	return func(options *Options) {
		options.MaxIdle = maxIdle
	}
}

func WithMaxActive(maxActive int) Option {
	return func(options *Options) {
		options.MaxActive = maxActive
	}
}

func WithIdleTimeout(idleTimeout time.Duration) Option {
	return func(options *Options) {
		options.IdleTimeout = idleTimeout
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(options *Options) {
		options.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(options *Options) {
		options.WriteTimeout = writeTimeout
	}
}

func WithConnectTimeout(connectTimeout time.Duration) Option {
	return func(options *Options) {
		options.ConnectTimeout = connectTimeout
	}
}

func (options *Options) init() *Options {
	return &Options{
		Host:           "127.0.0.1",
		Port:           6379,
		Password:       "",
		Database:       0,
		NetWork:        "tcp",
		MaxIdle:        10,
		MaxActive:      0,
		IdleTimeout:    time.Hour,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		ConnectTimeout: time.Second,
	}
}

func (options *LockOptions) init() *LockOptions {
	return &LockOptions{
		Expiry:     0,
		RetryDelay: 0,
		Tries:      0,
	}
}
