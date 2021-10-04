package errors

const (
	CodeOk Code = 200

	CodeValidateErr = 410

	CodeInnerError  Code = 500
	CodeServiceFail Code = 501
)

// Code 代表系统的唯一错误码
type Code uint

// Template 错误码对应的提示信息，该字符串有可能是一个模板字符串
func (c Code) Template() string {
	switch c {
	case CodeOk:
		return "请求成功"
	case CodeInnerError:
		return "服务器内部错误"
	case CodeValidateErr:
		return "校验出错"
	case CodeServiceFail:
		return "交易出错"
	}
	return ""
}
