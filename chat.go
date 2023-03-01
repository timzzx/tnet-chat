package main

import (
	"chat/internal/config"
	"chat/internal/handler"
	"chat/internal/svc"
	"chat/varx"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/timzzx/tnet"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := tnet.NewServer()
	cs := make(chan os.Signal)
	signal.Notify(cs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		select {
		case <-cs:
			{
				s.Stop() // 服务器退出
				os.Exit(1)
			}
		}
	}()

	// 增加路由
	s.AddHandlers(varx.LOGIN, handler.LoginHandler(varx.LOGIN, ctx))
	s.AddHandlers(varx.ROOMADD, handler.RoomAddHandler(varx.ROOMADD, ctx))
	// 服务启动
	s.Start()
}
