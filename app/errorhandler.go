package app

import (
	"git-knowledge/result"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (ehf *ErrorHandler) Handler(err error, translator *ut.Translator) *result.Response {
	if err == nil {
		panic("err不能为nil")
	}

	var response *result.Response

	switch err.(type) {

	case result.ServiceError: // 如果是服务错误，使用服务码构建返回对象
		e := err.(result.ServiceError)
		response = result.Build(e.Code).WithDetail(e.Detail)

	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		var errList []string
		for _, e := range errs {
			errList = append(errList, e.Translate(*translator))
		}
		response = result.Build(result.CodeReqParamErr).WithDetail(strings.Join(errList, "|"))

	default: // 如果是未知异常，则抛出系统内部错误
		response = result.Build(result.CodeInnerError).WithDetail(err.Error())
	}
	return response
}
