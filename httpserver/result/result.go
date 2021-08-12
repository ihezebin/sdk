package result

const (
	CodeBoolOk   = true
	CodeBoolFail = false
)

const (
	OK           = "OK"
	Fail         = "OK"
	Created      = "Created"
	Accepted     = "Accepted"
	NoContent    = "No Content"
	ResetContent = "Reset Content"

	BadRequest   = "Bad Request"
	Unauthorized = "Unauthorized"
	Forbidden    = "Forbidden"
	NotFound     = "Not Found"

	InternalServerError   = "Internal Server Error"
	InternalServerTimeout = "Internal Server Processing Timeout"
)

const (
	MessageOK           = "请求成功"
	MessageCreated      = "已成功创建资源"
	MessageAccepted     = "已经接受请求, 处理中"
	MessageNoContent    = "处理成功, 无其他响应信息"
	MessageResetContent = "已成功重置数据"

	MessageFail         = "请求失败"
	MessageBadRequest   = "请求的语法错误，服务器无法理解"
	MessageUnauthorized = "身份未认证"
	MessageForbidden    = "没有权限, 请求被服务器拒绝了"
	MessageNotFound     = "所请求的资源不存在"

	MessageInternalServerError   = "服务器内部错误, 无法完成请求"
	MessageInternalServerTimeout = "服务器处理超时"
)
