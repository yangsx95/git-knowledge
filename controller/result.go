package controller

import (
	"encoding/json"
	"fmt"
	"git-knowledge/errors"
)

type Response struct {
	Code errors.Code `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

func build(code errors.Code, fields ...interface{}) Response {
	target := code.Template()
	if fields != nil {
		target = fmt.Sprintf(target, fields)
	}
	return Response{
		Code: code,
		Msg:  target,
	}
}

var (
	Ok  = build(errors.CodeOk)
	Err = build(errors.CodeValidateErr)

	ErrParam = build(501, "参数有误")
)

// SetMsg 自定义响应信息
func (res *Response) SetMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// WithData 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		//Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
