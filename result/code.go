package result

// Code 代表系统的唯一错误码
type Code uint

// Code 枚举常量
const (
	CodeOk Code = 200

	CodeParseBodyErr = 400
	CodeNotFoundErr  = 404
	CodeMethodErr    = 405
	CodeValidateErr  = 410

	CodeInnerError  Code = 500
	CodeServiceFail Code = 501
)

// String 错误码基本描述
func (c Code) String() string {
	switch c {
	case CodeOk:
		return "请求成功"

	case CodeParseBodyErr:
		return "错误的请求内容，可能与Content-Type不匹配"
	case CodeNotFoundErr:
		return "接口不存在"
	case CodeValidateErr:
		return "请求参数有误"
	case CodeMethodErr:
		return "不支持的请求方法"

	case CodeServiceFail:
		return "交易出错"
	case CodeInnerError:
		return "服务器内部错误"
	}
	return ""
}
