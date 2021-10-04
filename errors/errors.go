package errors

import "fmt"

// 自定义项目异常

type ServiceError struct {
	Code Code
	Msg  string
}

// Error 实现error接口
func (e ServiceError) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.Code, e.Msg)
}

// New 创建一个Error
func New(code Code, fields ...interface{}) error {
	target := fmt.Sprintf(code.Template(), fields)
	return ServiceError{
		Code: code,
		Msg:  target,
	}
}
