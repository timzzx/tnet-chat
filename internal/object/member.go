package object

import "github.com/timzzx/tnet/types"

type Member struct {
	UserId int
	Name   string
	Agent  types.Connection
}

func NewMember(uid int, name string, agent types.Connection) *Member {
	return &Member{
		UserId: uid,
		Name:   name,
		Agent:  agent,
	}
}
