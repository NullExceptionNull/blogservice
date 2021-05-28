package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签失败")
	ErrorCreateTagtFail = NewError(20010002, "创建标签失败")
)
