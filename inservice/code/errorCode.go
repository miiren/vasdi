package code

var (
	CodeOk        = NewErrorCode("000000", "ok")
	CodeParamsErr = NewErrorCode("201001", "缺失必要参数")
	CodeSysErr    = NewErrorCode("201999", "系统其他错误")
)

type ErrorCode struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

func NewErrorCode(code string, msg string) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Msg:  msg,
	}
}
