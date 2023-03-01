package directives

import (
	"chat/jsontype"
	"chat/singleTest/types"
	"encoding/json"
	"fmt"
)

type RoomAdd struct {
}

func (h *RoomAdd) Do(agent types.Agent) error {
	if agent.GetUserId() == 0 {
		fmt.Println("请先登录")
		return nil
	}
	// 加入房间
	req, _ := json.Marshal(jsontype.RoomAddReq{
		UserId: agent.GetUserId(),
	})
	if err := agent.Send(2, req); err != nil {
		return err
	}
	msg := <-agent.GetResp()
	fmt.Println("==Resp:" + string(msg))

	return nil
}
