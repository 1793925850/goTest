package errcode

// 公共错误码

var (
	Success          = NewError(0, "成功")
	Fail             = NewError(100000000, "内部错误")
	InvalidParams    = NewError(100000001, "无效参数")
	Unauthorized     = NewError(100000002, "认证错误")
	NotFound         = NewError(100000003, "没有找到")
	Unknown          = NewError(100000004, "未知")
	DeadlineExceeded = NewError(100000005, "超出最后截止期限")
	AccessDenied     = NewError(100000006, "访问被拒绝")
	LimitExceed      = NewError(100000007, "访问限制")
	MethodNotAllowed = NewError(100000008, "不支持该方法")
)
