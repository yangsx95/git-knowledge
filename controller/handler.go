package controller

import (
	"git-knowledge/errors"
	"github.com/gin-gonic/gin"
)

// JSONHandler JSON处理器，将返回的数据以及错误信息解析为json
func JSONHandler(handler func(ctx *gin.Context) (interface{}, error)) func(ctx *gin.Context) {
	if handler == nil {
		panic("handler处理器不可以为空")
	}
	var response Response
	return func(ctx *gin.Context) {
		result, err := handler(ctx)
		if err != nil {
			response = parseErrToResp(err)
		} else {
			response = build(errors.CodeOk)
			response.Data = result
		}
		ctx.JSON(200, response)
	}
}

func parseErrToResp(err error) Response {
	if err == nil {
		panic("err不能为nil")
	}
	var response Response
	switch err.(type) {
	case errors.ServiceError:
		e := err.(errors.ServiceError)
		response = build(e.Code, e.Msg)
	default: // 如果是未知异常，则抛出系统内部错误
		response = build(errors.CodeInnerError)
	}
	return response
}
