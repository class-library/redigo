package redigo

import "github.com/gomodule/redigo/redis"

// Cmd 封装后的redis执行命令
func Cmd(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, key)
	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}
	return Exec(cmd, params...)
}

// Exec redigo原生执行命令
func Exec(cmd string, params ...interface{}) (interface{}, error) {
	con := Redis.Pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	return con.Do(cmd, params...)
}

// Set set命令
func Set(key string, val interface{}) (interface{}, error) {
	return Cmd("set", key, val)
}

// Get get命令
func Get(key string) (val string, err error) {
	var exist interface{}
	if exist, err = Cmd("get", key); err != nil {
		return
	}
	if exist != nil {
		return redis.String(exist, nil)
	}
	return
}

// SetEx setEx命令
func SetEx(key string, ttl int64, val interface{}) (interface{}, error) {
	return Cmd("setEx", key, ttl, val)
}

// Expire expire命令[给key设定一个存活时间（单位为秒）]
func Expire(key string, ttl int64) (bool, error) {
	return redis.Bool(Cmd("expire", key, ttl))
}

// ExpireAt expireAt命令[以UNIX时间戳格式设置key的过期时间]
func ExpireAt(key string, timestamp int64) (bool, error) {
	return redis.Bool(Cmd("expireAt", key, timestamp))
}

// Del del命令
func Del(key string) (bool, error) {
	return redis.Bool(Cmd("del", key))
}

// Ttl ttl命令
func Ttl(key string) (int64, error) {
	return redis.Int64(Cmd("ttl", key))
}

// HMSet hMSet命令
func HMSet(key string, val interface{}) (interface{}, error) {
	return Exec("HMSet", redis.Args{}.Add(key).AddFlat(val)...)
}

// HGetAll hGetAll命令
func HGetAll(key string) ([]interface{}, error) {
	return redis.Values(Exec("HGetAll", key))
}

// Keys keys命令
func Keys(keys string) ([]string, error) {
	return redis.Strings(Exec("keys", keys))
}

// Persist persist命令[移除给定key的过期时间]
func Persist(key string) (bool, error) {
	return redis.Bool(Cmd("persist", key))
}

// Incr 自增+1
func Incr(key string) (int64, error) {
	return redis.Int64(Exec("incr", key))
}

// IncrBy 自增+x
func IncrBy(key string, val int64) (int64, error) {
	return redis.Int64(Exec("incrBy", key, val))
}

// Decr 自减-1
func Decr(key string) (int64, error) {
	return redis.Int64(Exec("decr", key))
}

// DecrBy 自减-x
func DecrBy(key string, val int64) (int64, error) {
	return redis.Int64(Exec("decrBy", key, val))
}

// Exists 检测KEY是否存在
func Exists(key string) (bool, error) {
	return redis.Bool(Exec("exists", key))
}
