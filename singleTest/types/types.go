package types

type Directives interface {
	Do(agent Agent) error
}

type Agent interface {
	Send(rid int, data []byte) error
	GetResp() <-chan []byte
	Stop()
}
