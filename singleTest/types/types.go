package types

type Directives interface {
	Do(agent Agent)
}

type Agent interface {
	Send(rid int, data []byte) error
}
