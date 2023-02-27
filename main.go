package main

import (
	handlers "chat/Handlers"

	"github.com/timzzx/tnet"
)

func main() {
	s := tnet.NewServer()
	// 添加一个handler
	s.AddHandlers(1, handlers.NewTestHandler(1))
	// 服务启动
	s.Start()
}
