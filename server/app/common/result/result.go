package result

// Result 响应结果
type HttpResult struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

var httpResult HttpResult

// Success 成功
func Success(data interface{}) HttpResult {
	httpResult.Code = 200
	httpResult.Msg = "success"
	httpResult.Data = data
	return httpResult
}

// Fail 失败
func Fail(msg string) HttpResult {
	httpResult.Code = 400
	httpResult.Msg = msg
	return httpResult
}

// FailWithCode 失败
func FailWithCode(code int, msg string) HttpResult {
	httpResult.Code = code
	httpResult.Msg = msg
	return httpResult
}

func Error() string {
	return httpResult.Msg
}

// 构造函数
func NewResult() HttpResult {
	return HttpResult{}
}
