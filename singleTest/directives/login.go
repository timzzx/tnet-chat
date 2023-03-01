package directives

import (
	"chat/jsontype"
	"chat/singleTest/types"
	"encoding/json"
	"fmt"
)

type Login struct {
}

func (h *Login) Do(agent types.Agent) {
	// 登录
	req, _ := json.Marshal(jsontype.LoginReq{
		Name:     "timzzx",
		Password: "123456",
	})
	if err := agent.Send(1, req); err != nil {
		fmt.Println("消息发送失败", err)
	}
}
