package result

type Message struct {
	Func         string `json:"func"`
	ContentType  string `json:"content_type"`
	Content      string `json:"content"`
	Success      bool   `json:"success"`       // 客户端发送的消息不存在该字段
	ErrorMessage string `json:"error_message"` // 客户端发送的消息不存在该字段
}

func NewEmptyMessage() *Message {
	return &Message{}
}

func NewErrorMessage(topic string, err error) *Message {
	return &Message{
		Func:         topic,
		ContentType:  "text/plain",
		Content:      "",
		Success:      false,
		ErrorMessage: err.Error(),
	}
}

func NewSuccessMessage(fun, contentType, content string) *Message {
	return &Message{
		Func:         fun,
		ContentType:  contentType,
		Content:      content,
		Success:      true,
		ErrorMessage: "",
	}
}
