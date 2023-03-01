package directives

import (
	"chat/jsontype"
	"chat/singleTest/types"
	"encoding/json"
	"fmt"
	"strconv"
)

type Login struct {
}

func (h *Login) Do(agent types.Agent) error {
	// 登录
	// req, _ := json.Marshal(jsontype.LoginReq{
	// 	Name:     "timzzx",
	// 	Password: "123456",
	// })
	req, _ := json.Marshal(jsontype.LoginReq{
		Name:     "ttt",
		Password: "123456",
	})
	if err := agent.Send(1, req); err != nil {
		return err
	}
	msg := <-agent.GetResp()
	fmt.Println("==Resp:" + string(msg))
	// 设置userid
	var data jsontype.LoginRsp
	json.Unmarshal(msg, &data)
	agent.SetUserId(data.UserId)
	fmt.Println("UserId:" + strconv.Itoa(agent.GetUserId()))

	return nil
}
