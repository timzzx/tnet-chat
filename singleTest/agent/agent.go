package agent

import (
	"chat/singleTest/directives"
	"chat/singleTest/types"
	"fmt"
	"net"

	"github.com/timzzx/tnet"
)

type agent struct {
	conn       net.Conn
	resp       chan []byte
	Directives map[string]types.Directives
}

func NewAgent(conn net.Conn) *agent {
	return &agent{
		conn:       conn,
		Directives: make(map[string]types.Directives),
		resp:       make(chan []byte),
	}
}

// 发送消息
func (a *agent) Send(rid int, data []byte) error {
	msg, err := tnet.Pack(rid, data)
	_, err = a.conn.Write(msg)
	return err
}

// 接收消息
func (a *agent) Reader() {
	defer a.Stop()
	for {
		// 接收消息
		rid, data, err := tnet.Unpack(a.conn)
		if err != nil {
			fmt.Println("消息收回", err)
			return
		}

		if rid != 0 {
			// fmt.Println("Resp:" + string(data))
			// fmt.Print("> ")
			a.resp <- data
		}
	}
}

// 获取消息
func (a *agent) GetResp() <-chan []byte {
	return a.resp
}

// 停止
func (a *agent) Stop() {
	a.conn.Close()
}

// 初始化指令
func (a *agent) InitDirectives() {
	a.Directives["login"] = &directives.Login{}
}
