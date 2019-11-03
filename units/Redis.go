package units

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"wb-impulse-sender/conf"
	"time"
)

//redis接口
type redis_i interface {
	//List——追加元素
	Push(key string, value interface{}) (err error)
	//Len——获取List长度
	Len(key string) (l int64, err error)
	//LPop——list头弹出
	LPop(key string) (ret string, err error)
}

//redis类
type redisRealize struct {
	cfg  *conf.Config
	key  string
	connect *redis.Pool
}

//redis初始化
func newRedis(cfg *conf.Config, key string) (r *redisRealize) {
	//连接
	connect := connector(cfg, key)
	return &redisRealize{
		cfg : cfg,
		key : key,
		connect : connect,
	}
}

//连接
func connector(cfg *conf.Config, key string) *redis.Pool {
	redisConfig := cfg.Redis[key]
	//fmt.Println(redisConfig, time.Duration(redisConfig.IdleTimeout * time.Second))
	pool := &redis.Pool{
		MaxIdle:     redisConfig.Idle,
		MaxActive:   redisConfig.Active,
		IdleTimeout: time.Duration(redisConfig.IdleTimeout * time.Second),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(redisConfig.Network, redisConfig.Addr,
				redis.DialConnectTimeout(time.Duration(redisConfig.DialTimeout * time.Millisecond)),
				redis.DialReadTimeout(time.Duration(redisConfig.ReadTimeout * time.Millisecond)),
				redis.DialWriteTimeout(time.Duration(redisConfig.WriteTimeout * time.Millisecond)),
				redis.DialPassword(redisConfig.Password),
				redis.DialDatabase(redisConfig.Db),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return pool
}

/**
	zset——添加filed
 */
func (r *redisRealize) ZAdd(key string, filed string, score interface{}) (err error) {
	conn := r.connect.Get()
	defer conn.Close()
	if err = conn.Send("ZADD", key, score, filed); err != nil {
		return err
	}
	if err := conn.Flush(); err != nil {
		return err
	}
	if _, err := conn.Receive(); err != nil {
		return err
	}
	return
}

/**
	List——追加元素
 */
func (r *redisRealize) Push(key string, value interface{}) (err error) {
	conn := r.connect.Get()
	defer conn.Close()
	if err = conn.Send("RPUSH", key, value); err != nil {
		fmt.Println(err)
		return
	}
	if err := conn.Flush(); err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := conn.Receive(); err != nil {
		fmt.Println(err)
		return err
	}
	return
}

/**
	Len——获取List长度
 */
func (r *redisRealize) Len(key string) (l int64, err error) {
	conn := r.connect.Get()
	defer conn.Close()
	if err = conn.Send("LLEN", key); err != nil {
		fmt.Println(err)
		return 0, err
	}
	if err = conn.Flush(); err != nil {
		fmt.Println(err)
		return 0, err
	}
	var length interface{}
	if length, err = conn.Receive(); err != nil {
		fmt.Println(length, err)
		return 0, err
	}
	return  length.(int64), err
}

/**
	LPop——list头弹出
 */
func (r *redisRealize) LPop(key string) (ret string, err error) {
	conn := r.connect.Get()
	defer conn.Close()
	res, err := conn.Do("LPOP", key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ret = fmt.Sprintf("%s", res)
	return  ret, err
}

