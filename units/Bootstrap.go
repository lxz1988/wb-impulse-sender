package units

import (
	"wb-impulse-sender/conf"
	"sync"
)

var once sync.Once

var instance units

//library结构体
type units struct {
	Redis redis_i //redis实现
}

/**
	获取units实例
 */
func GetUnitsInstance(cfg *conf.Config, key string) (*units) {
	once.Do(func() {
		instance.Redis = newRedis(cfg, key)
	})
	return &instance
}
