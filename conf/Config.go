package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Desc 	string
	Server 	serverConfig
	Log 	logConfig
	Mysql 	map[string]mysqlConfig
	Redis 	map[string]redisConfig
	Number 	numberConfig
}

//服务配置信息
type serverConfig struct {
	Desc string
}
//日志配置信息
type logConfig struct {
	Desc 		string
	DumpPeriod 	time.Duration `toml:"dump_period"`
}

//Mysql配置信息
type mysqlConfig struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   string
}

//Redis配置信息
type redisConfig struct {
	Network      string	 `toml:"network"`
	Addr         string	 `toml:"addr"`
	Db			 int
	Password     string
	Active       int	//最大连接数 默认0 不限制
	Idle         int	//最大的空闲连接数
	DialTimeout  time.Duration	//最大连接超时时间 毫秒
	ReadTimeout  time.Duration	//最大读取超时时间 毫秒
	WriteTimeout time.Duration	//同上 毫秒
	IdleTimeout  time.Duration	//连接最大空闲超时时间 秒
	Expire       time.Duration	//过期时间 秒
}

type numberConfig struct {
	RedisKey 		string 				`toml:"redis_key"` 			//存储list的redis
	ListKey 		string 				`toml:"list_key"` 			//存储号码的list key
	Threshold		int64		 		`toml:"threshold"`			//每个list中最大数量
	GeneratePeriod	time.Duration		`toml:"generate_period"`	//生成号码的周期，毫秒
}

var cfg Config

var instance sync.Once

/**
	获取配置文件
 */
func GetConfigInstance() (config *Config) {
	instance.Do(func() {
		config_file := flag.String("conf","config/app.dev.conf","config file")
		flag.Parse()
		_,err := toml.DecodeFile(*config_file,&cfg)
		if err != nil {
			msg := strings.Join([]string{"load config file error:",err.Error()},"")
			panic(msg)
		}
	})
	return &cfg
}
