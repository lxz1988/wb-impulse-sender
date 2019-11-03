package main

import (
	"fmt"
	"os"
	"os/signal"
	"wb-impulse-sender/bootstrap"
	"wb-impulse-sender/conf"
	"wb-impulse-sender/library"
	"wb-impulse-sender/pond"
	"strings"
	"syscall"
	"time"
)

func main()  {
	fmt.Println("wb_impulse_sender start...")

	sys_chan := make(chan os.Signal, 1)
	signal.Notify(sys_chan, syscall.SIGINT, syscall.SIGTERM)

	//获取配置信息
	cfg := conf.GetConfigInstance()
	//初始化服务
	server, err := pond.NewServer(cfg)

	//units.GetUnitsInstance(cfg, "pond").Redis.Push("zfj","zhangfj")

	if err != nil {
		msg := strings.Join([]string{"wb_impulse_sender server new error : ", err.Error()},"")
		panic(msg)
	}
	if err := server.Start(); err != nil {
		msg := strings.Join([]string{"wb_impulse_sender server start error : ", err.Error()},"")
		panic(msg)
	}
	//启动Http
	bootstrap.Boot()

	ticker := time.NewTicker(cfg.Log.DumpPeriod * time.Second)
	for {
		select {
		case <- ticker.C:
			library.GetLibraryInstance().Logger.Log("wb_impulse_sender heartbeat...")
		case <- sys_chan:
			if err := server.Stop(); err != nil {
				msg := strings.Join([]string{"wb_impulse_sender server stop error : ", err.Error()},"")
				library.GetLibraryInstance().Logger.Log(msg)
			}else{
				library.GetLibraryInstance().Logger.Log("wb_impulse_sender server stop...")
			}
			return
		}
	}
}
