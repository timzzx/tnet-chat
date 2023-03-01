package directives

import (
	"chat/jsontype"
	"chat/singleTest/types"
	"encoding/json"
	"fmt"
)

type Login struct {
}

func (h *Login) Do(agent types.Agent) error {
	// 登录
	req, _ := json.Marshal(jsontype.LoginReq{
		Name:     "timzzx",
		Password: "123456",
	})
	if err := agent.Send(122, req); err != nil {
		return err
	}
	msg := <-agent.GetResp()
	fmt.Println(">Resp:" + string(msg))
	return nil
}
