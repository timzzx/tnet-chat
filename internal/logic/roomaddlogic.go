package logic

import (
	"chat/internal/object"
	"chat/internal/svc"
	"chat/jsontype"
	"encoding/json"

	"github.com/timzzx/tnet"
	"github.com/timzzx/tnet/types"
)

type RoomAddLogic struct {
	id     int
	svcCtx *svc.ServiceContext
}

func NewRoomAddLogic(id int, svcCtx *svc.ServiceContext) types.Handler {
	return &RoomAddLogic{
		id:     id,
		svcCtx: svcCtx,
	}
}

func (l *RoomAddLogic) Do(req []byte, agent types.Connection) {
	// 请求参数
	var Req jsontype.RoomAddReq
	if err := json.Unmarshal(req, &Req); err != nil {
		resp, _ := json.Marshal(jsontype.RoomAddResp{
			Code: 500,
			Msg:  err.Error(),
		})
		msg, _ := tnet.Pack(l.id, resp)
		agent.Send(msg)
		return
	}
	// 从数据库获取用户信息
	u := l.svcCtx.ChatModel.User
	user, err := u.WithContext(agent.Ctx()).Where(u.ID.Eq(int64(Req.UserId))).First()
	if err != nil {
		resp, _ := json.Marshal(jsontype.RoomAddResp{
			Code: 500,
			Msg:  err.Error(),
		})
		msg, _ := tnet.Pack(l.id, resp)
		agent.Send(msg)
		return
	}

	m := object.NewMember(int(user.ID), user.Name, agent)
	object.Room.Add(m)

	// 发消息给自己
	resp, _ := json.Marshal(jsontype.RoomAddResp{
		Code: 200,
		Msg:  "加入成功",
	})
	msg, _ := tnet.Pack(l.id, resp)
	agent.Send(msg)

	// 发消息给所有人
	resp, _ = json.Marshal(jsontype.RoomAddResp{
		Code: 200,
		Msg:  user.Name + "加入成功",
	})
	msg, _ = tnet.Pack(1002, resp)
	for _, member := range object.Room.List() {
		member.Agent.Send(msg)
	}

	return
}
