package result

import "fmt"

// ServiceError 系统服务错误
type ServiceError struct {
	Code   Code   // 错误编码
	Detail string // 错误详细信息
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("code:%d,detail:%v", e.Code, e.Detail)
}

// ErrorOf 创建一个 ServiceError
func ErrorOf(code Code) error {
	return ServiceError{
		Code: code,
	}
}

func ErrorOfWithDetail(code Code, detail string) error {
	return ServiceError{
		Code:   code,
		Detail: detail,
	}
}

func ErrorOfWithDetailF(code Code, detailTemplate string, params ...interface{}) error {
	return ErrorOfWithDetail(code, fmt.Sprintf(detailTemplate, params...))
}
