package result

// Code 代表系统的唯一错误码
type Code uint

// Code 枚举常量
const (
	CodeOk Code = 200

	CodeValidateErr = 410
	CodeReqParamErr = 411

	CodeInnerError  Code = 500
	CodeServiceFail Code = 501
)

// String 错误码基本描述
func (c Code) String() string {
	switch c {
	case CodeOk:
		return "请求成功"

	case CodeValidateErr:
		return "校验出错"
	case CodeReqParamErr:
		return "请求参数有误"

	case CodeServiceFail:
		return "交易出错"
	case CodeInnerError:
		return "服务器内部错误"
	}
	return ""
}
