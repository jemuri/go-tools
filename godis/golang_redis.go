package godis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang.guazi-corp.com/finance/aegis-common/utils/config"
	"sync"
)

const (
	NetWorkDefault  = "tcp"
	AddressDefault  = "127.0.0.1:6379"
	PasswordDefault = ""
)

var onceDo sync.Once
var pool *redis.Pool

type RedisPool struct {
	Pool *redis.Pool
}

// NewRedisPool NewRedisPool
func NewRedisPool() *RedisPool {
	onceDo.Do(func() {
		pool = &redis.Pool{
			MaxIdle:   3, /*最大的空闲连接数*/
			MaxActive: 8, /*最大的激活连接数*/
			Dial:      DialHandler,
		}
	})

	return &RedisPool{
		Pool: pool,
	}
}

func DialHandler() (redis.Conn, error) {
	network := config.CertainString("redis/network")
	if network == "" {
		network = NetWorkDefault
	}

	address := config.CertainString("redis/address")
	if address == "" {
		address = AddressDefault
	}

	password := config.CertainString("redis/password")
	if password == "" {
		password = PasswordDefault
	}

	c, err := redis.Dial(network, address, redis.DialPassword(password))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *RedisPool) Get(key string) (string, error) {
	c := r.Pool.Get()
	defer closeConn(c)

	return redis.String(c.Do("GET", key))
}

func (r *RedisPool) GetInt64(key string) (int64, error) {
	c:= r.Pool.Get()
	defer  closeConn(c)

	return redis.Int64(c.Do("GET",key))
}

//在SET命令中，有很多选项可用来修改命令的行为。 以下是SET命令可用选项的基本语法
//SET KEY VALUE [EX seconds] [PX milliseconds] [NX|XX]
// EX seconds − 设置指定的到期时间(以秒为单位)。
// PX milliseconds - 设置指定的到期时间(以毫秒为单位)。
// NX - 仅在键不存在时设置键。
// XX - 只有在键已存在时才设置。
// SET mykey "redis" EX 60 NX
// 以上示例将在键“mykey”不存在时，设置键的值，到期时间为60秒
func (r *RedisPool) Set(key string, value interface{}) error {
	c := r.Pool.Get()
	defer closeConn(c)

	_, err := c.Do("SET", key, value)
	return err
}

func (r *RedisPool) SetEX(key string, value interface{}, seconds int64) error {
	c := r.Pool.Get()
	defer closeConn(c)

	_, err := c.Do("SET", key, value, "EX", seconds)

	return err
}
func (r *RedisPool) SetNX(key string, value interface{}, seconds int64) error {
	c := r.Pool.Get()
	defer closeConn(c)

	var err error
	if seconds <= 0 {
		_, err = c.Do("SET", key, value, "NX")
	} else {
		_, err = c.Do("SET", key, value, "EX", seconds, "NX")
	}

	return err
}

func (r *RedisPool) Del(key string) error {
	c := r.Pool.Get()
	defer closeConn(c)

	_, err := c.Do("DEL", key)
	return err
}

func closeConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("redis.Conn close fail! err: ", err)
	}
}
