package defs

type Err struct {
	Error     string `json: "error"`
	ErrorCode string `json: "error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "请求体不存在", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "用户验证不通过", ErrorCode: "002"}}
)
