package jsontype

// 登录请求
type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 登录返回
type LoginRsp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
