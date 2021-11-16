package v1

import (
	"fmt"
	"git-knowledge/api/v1/vo"
)

type MessageApi interface {
	SayHello(ws vo.WebsocketSender, content *string) error
}

type messageApiImpl struct {
}

func NewMessageApi() MessageApi {
	return &messageApiImpl{}
}

func (r *messageApiImpl) SayHello(ws vo.WebsocketSender, content *string) error {
	fmt.Println("收到客户端消息：" + *content)
	return ws.Send("你好呀客户端")
}
