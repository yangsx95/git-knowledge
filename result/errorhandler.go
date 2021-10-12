package result

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/url"
	"strings"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (ehf *ErrorHandler) Handler(err error, translator *ut.Translator) *Response {
	if err == nil {
		panic("err不能为nil")
	}

	var response *Response

	switch err.(type) {

	case ServiceError: // 如果是服务错误，使用服务码构建返回对象
		e := err.(ServiceError)
		response = Build(e.Code).WithDetail(e.Detail)

	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		var errList []string
		for _, e := range errs {
			errList = append(errList, e.Translate(*translator))
		}
		response = Build(CodeValidateErr).WithDetail(strings.Join(errList, "|"))

	case *url.Error:
		e := err.(*url.Error)
		u, _ := url.Parse(e.URL)
		switch u.Host {
		case "github.com":
			response = Build(CodeGithubConnectionErr).WithDetail(e.Error())
		default:
			response = Build(CodeInnerError).WithDetail(err.Error())
		}
	case *echo.HTTPError:
		err := err.(*echo.HTTPError)
		switch err.Code {
		case 400:
			response = Build(CodeParseBodyErr)
		case 404:
			response = Build(CodeNotFoundErr)
		case 405:
			response = Build(CodeMethodErr)
		default:
			response = Build(CodeInnerError).WithDetail(err.Error())
		}
	default: // 如果是未知异常，则抛出系统内部错误
		response = Build(CodeInnerError).WithDetail(err.Error())
	}
	return response
}
