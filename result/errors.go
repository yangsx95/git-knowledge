package result

import "fmt"

// ServiceError 系统服务错误
type ServiceError struct {
	Code Code   // 错误编码
	Msg  string // 错误信息
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("code:%d,detail:%v", e.Code, e.Msg)
}

// ErrorOf 创建一个 ServiceError
func ErrorOf(code Code) error {
	return ServiceError{
		Code: code,
		Msg:  code.String(),
	}
}
