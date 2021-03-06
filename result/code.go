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
	CodeAuthErr      = 420

	CodeUserNotExists               = 430
	CodeWrongPassword               = 431
	CodeRegisterUserIdAlreadyExists = 440
	CodeRegisterEmailAlreadyExists  = 441
	CodeGithubAuthFail              = 450
	CodeCredentialUseless           = 460

	CodeInnerError  Code = 500
	CodeServiceFail Code = 501

	CodeGithubConnectionErr Code = 571
	CodeGithubLoginErr      Code = 572
)

// String 错误码基本描述
func (c Code) String() string {
	switch c {
	case CodeOk:
		return "请求成功"

	case CodeParseBodyErr:
		return "错误的请求内容，可能与Content-Type不匹配或者参数有误"
	case CodeNotFoundErr:
		return "接口不存在"
	case CodeValidateErr:
		return "请求参数有误"
	case CodeAuthErr:
		return "认证失败，你没有权限访问"
	case CodeMethodErr:
		return "不支持的请求方法"
	case CodeUserNotExists:
		return "用户不存在"
	case CodeWrongPassword:
		return "密码错误"
	case CodeRegisterUserIdAlreadyExists:
		return "用户Id已经存在"
	case CodeRegisterEmailAlreadyExists:
		return "邮箱已经被注册"
	case CodeGithubAuthFail:
		return "Github登录授权失败"

	case CodeCredentialUseless:
		return "无效的凭证"

	case CodeServiceFail:
		return "交易出错"
	case CodeInnerError:
		return "服务器内部错误"

	case CodeGithubConnectionErr:
		return "与github服务器建立连接失败"
	case CodeGithubLoginErr:
		return "Github登录失败"
	}
	return ""
}
