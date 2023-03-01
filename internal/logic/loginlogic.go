package logic

import (
	"chat/internal/svc"
	"chat/jsontype"
	"encoding/json"
	"fmt"

	"github.com/timzzx/tnet"
	"github.com/timzzx/tnet/types"
)

type LoginLogic struct {
	id     int
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(id int, svcCtx *svc.ServiceContext) types.Handler {
	return &LoginLogic{
		id:     id,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Do(req []byte, agent types.Connection) {
	// 获取参数
	var Req jsontype.LoginReq
	if err := json.Unmarshal(req, &Req); err != nil {
		resp, _ := json.Marshal(jsontype.LoginRsp{
			Code: 500,
			Msg:  err.Error(),
		})
		msg, _ := tnet.Pack(l.id, resp)
		agent.Send(msg)
		return
	}
	// 数据库获取
	u := l.svcCtx.ChatModel.User
	user, err := u.WithContext(agent.Ctx()).Where(u.Name.Eq(Req.Name)).Debug().First()
	if err != nil {
		resp, _ := json.Marshal(jsontype.LoginRsp{
			Code: 500,
			Msg:  err.Error(),
		})
		msg, _ := tnet.Pack(l.id, resp)
		agent.Send(msg)
		return
	}

	// 密码错误
	if user.Password != Req.Password {
		fmt.Println(user.Password)
		resp, _ := json.Marshal(jsontype.LoginRsp{
			Code: 500,
			Msg:  "密码错误",
		})
		msg, _ := tnet.Pack(l.id, resp)
		agent.Send(msg)
		return
	}
	resp, _ := json.Marshal(jsontype.LoginRsp{
		UserId: int(user.ID),
		Code:   200,
		Msg:    "登录成功",
	})
	msg, _ := tnet.Pack(l.id, resp)
	agent.Send(msg)

	// resp, _ = json.Marshal(jsontype.LoginRsp{
	// 	Code: 200,
	// 	Msg:  user.Name + "登录成功",
	// })
	// msg, _ = tnet.Pack(1001, resp)
	// agent.Send(msg)
	return
}
