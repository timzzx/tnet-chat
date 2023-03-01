package jsontype

type RoomAddReq struct {
	UserId int `json:"user_id"`
}

type RoomAddResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
