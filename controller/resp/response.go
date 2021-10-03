package resp

import "encoding/json"

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

func build(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

var (
	Ok  = build(200, "success")
	Err = build(500, "系统内部错误")

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
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
