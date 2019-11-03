package pond

import (
	"wb-impulse-sender/conf"
	"wb-impulse-sender/library"
	"wb-impulse-sender/units"
	"time"
)

//服务类
type Server struct {
	//配置信息
	cfg *conf.Config
}

//服务初始化
func NewServer(cfg *conf.Config) (s *Server, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = library.GetLibraryInstance().Exception.Error(er)
		}
	}()
	s = &Server{
		cfg: cfg,
	}
	return s, nil
}

//启动服务
func (s *Server) Start() error {
	//获取NodeID
	wn := NewWorkNode(s.cfg, "pond")
	go generate(wn, s.cfg)
	return nil
}

//生成号码
func generate(wn *WorkNode, cfg *conf.Config)  {
	//每隔N毫秒 开始判断list中号码的数量
	//当小于阈值时开始生成号码
	ticker := time.NewTicker(cfg.Number.GeneratePeriod * time.Millisecond)

	for {
		select {
		case <- ticker.C:
			//获取list中号码个数
			numberCount,_ := units.GetUnitsInstance(cfg, cfg.Number.RedisKey).Redis.Len(cfg.Number.ListKey)
			//fmt.Println(numberCount)
			//当list中号码小于一定阈值时，开始预置号码
			for numberCount < cfg.Number.Threshold {
				//生成number
				number := wn.Generate()
				//写入redis
				units.GetUnitsInstance(cfg, cfg.Number.RedisKey).Redis.Push(cfg.Number.ListKey, number)
				numberCount++
			}
		}

	}
}

//停止服务
func (s *Server) Stop() error {
	return nil
}