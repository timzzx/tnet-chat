package handler

import (
	"chat/internal/logic"
	"chat/internal/svc"

	"github.com/timzzx/tnet/types"
)

func RoomAddHandler(id int, svc *svc.ServiceContext) types.Handler {
	return logic.NewRoomAddLogic(id, svc)
}
