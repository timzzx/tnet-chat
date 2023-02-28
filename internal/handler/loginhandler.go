package handler

import (
	"chat/internal/logic"
	"chat/internal/svc"

	"github.com/timzzx/tnet/types"
)

func LoginHandler(id int, svc *svc.ServiceContext) types.Handler {
	return logic.NewLoginLogic(id, svc)
}
