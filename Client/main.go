package main

import (
	"chat/jsontype"
	"encoding/json"
	"fmt"
	"net"

	"github.com/timzzx/tnet"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.13:9999")
	if err != nil {
		fmt.Println("连接失败", err)
	}
	defer conn.Close()

	// 发送错误结构的消息
	msg, err := tnet.Pack(1, []byte("test"))
	conn.Write(msg)
	if err != nil {
		fmt.Println("消息发送失败", err)
		return
	}
	// 接收消息
	_, data, err := tnet.Unpack(conn)
	if err != nil {
		fmt.Println("消息收回", err)
		conn.Close()
		return
	}

	fmt.Println(string(data))
	// 发送密码错误的消息
	req, _ := json.Marshal(jsontype.LoginReq{
		Name:     "timzzx",
		Password: "123",
	})
	msg, err = tnet.Pack(1, req)
	conn.Write(msg)
	if err != nil {
		fmt.Println("消息发送失败", err)
		return
	}
	// 接收消息
	_, data, err = tnet.Unpack(conn)
	if err != nil {
		fmt.Println("消息收回", err)
		conn.Close()
		return
	}

	fmt.Println(string(data))
	// 发送正确的消息
	req, _ = json.Marshal(jsontype.LoginReq{
		Name:     "timzzx",
		Password: "123456",
	})
	msg, err = tnet.Pack(1, req)
	conn.Write(msg)
	if err != nil {
		fmt.Println("消息发送失败", err)
		return
	}

	// 接收消息
	_, data, err = tnet.Unpack(conn)
	if err != nil {
		fmt.Println("消息收回", err)
		conn.Close()
		return
	}

	fmt.Println(string(data))

}
