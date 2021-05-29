package errcode

var (
	OrderParamsError = NewError(3002, "请求参数解析错误")
	TokenError       = NewError(3001, "Token参数错误")
)
