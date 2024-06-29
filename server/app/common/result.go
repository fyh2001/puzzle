package common

// Result 响应结果
type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

var HttpResult Result

// Success 成功
func (Result) Success(data interface{}) Result {
	HttpResult.Code = 200
	HttpResult.Msg = "success"
	HttpResult.Data = data
	return HttpResult
}

// Fail 失败
func (Result) Fail(msg string) Result {
	HttpResult.Code = 400
	HttpResult.Msg = msg
	return HttpResult
}

// FailWithCode 失败
func (Result) FailWithCode(code int, msg string) Result {
	HttpResult.Code = code
	HttpResult.Msg = msg
	return HttpResult
}

func (Result) Error() string {
	return HttpResult.Msg
}
