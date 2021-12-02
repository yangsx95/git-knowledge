package vo

import (
	"git-knowledge/result"
	"git-knowledge/ws"
)

type LoginInfo struct {
	// UserId 当前登录的用户id
	Userid string `json:"userid"`
}

// WebsocketSender websocket发送器
type WebsocketSender struct {
	Conn        *ws.Connection // websocket连接对象
	Func        string         // 接收到的当前消息的主题，是由接受的客户端消息决定（使用值拷贝的方式传递给使用者，防止并发安全）
	ContentType string         // 接收到的当前消息的内容类型，接受的内容类型为json，返回的类型必须也是json
}

func (sender *WebsocketSender) Send(value interface{}) error {
	s := result.GetSerializerByContentType(sender.ContentType)
	v, err := s.Serialize(value)
	if err != nil {
		return err
	}
	m := result.NewSuccessMessage(sender.Func, sender.ContentType, string(v))
	return sender.Conn.SendJSON(m)
}
