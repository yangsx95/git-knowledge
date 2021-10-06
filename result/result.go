package result

// Response 响应体
type Response struct {
	Code   int         `json:"code"`   // 错误码
	Msg    string      `json:"msg"`    // 错误描述
	Detail string      `json:"detail"` // 详细描述
	Data   interface{} `json:"data"`   // 返回数据
}

// WithDetail 追加描述信息
func (res *Response) WithDetail(detail string) *Response {
	res.Detail = detail
	return res
}

// WithData 追加响应数据
func (res *Response) WithData(data interface{}) *Response {
	res.Data = data
	return res
}

// Build 构造结果对象
func Build(code Code) *Response {
	return &Response{
		Code: int(code),
		Msg:  code.String(),
	}
}
