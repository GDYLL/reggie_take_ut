package common

// R 返回结果结构体
type R struct {
	Code int         `json:"code"` // 编码：1成功，0和其它数字为失败
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 数据
}

type Result struct{}

// Success 返回成功结果
func (r Result) Success(data interface{}) *R {
	return &R{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
}

// Error 返回失败结果
func (r Result) Error(msg string) *R {
	return &R{
		Code: 0,
		Msg:  msg,
	}
}
