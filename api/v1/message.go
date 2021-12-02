package v1

import (
	"fmt"
	"git-knowledge/api/v1/vo"
)

type p struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MessageApi interface {
	SayText(ws vo.WebsocketSender, content string) error
	SayJSON(ws vo.WebsocketSender, content p) error
	SayJSONArray(ws vo.WebsocketSender, content []string) error
	SayPtrText(ws vo.WebsocketSender, content *string) error
	SayPtrJSON(ws vo.WebsocketSender, content *p) error
	SayPtrJSONArray(ws vo.WebsocketSender, content *[]string) error
}

type messageApiImpl struct {
}

func NewMessageApi() MessageApi {
	return &messageApiImpl{}
}

func (r *messageApiImpl) SayText(ws vo.WebsocketSender, content string) error {
	fmt.Println("收到客户端消息：" + content)
	return ws.Send("你好呀客户端")
}

func (r *messageApiImpl) SayJSON(ws vo.WebsocketSender, content p) error {
	fmt.Printf("收到客户端消息：%v, %v\n", content.Name, content.Age)
	return ws.Send(&content)
}

func (r *messageApiImpl) SayJSONArray(ws vo.WebsocketSender, content []string) error {
	fmt.Printf("收到客户端消息：%v\n", content)
	return ws.Send(&[]string{"你好呀", "客户端"})
}

func (r *messageApiImpl) SayPtrText(ws vo.WebsocketSender, content *string) error {
	fmt.Println("收到客户端消息：" + *content)
	return ws.Send("你好呀客户端")
}

func (r *messageApiImpl) SayPtrJSON(ws vo.WebsocketSender, content *p) error {
	fmt.Printf("收到客户端消息：%v, %v\n", content.Name, content.Age)
	return ws.Send(content)
}

func (r *messageApiImpl) SayPtrJSONArray(ws vo.WebsocketSender, content *[]string) error {
	fmt.Printf("收到客户端消息：%v\n", *content)
	return ws.Send(&[]string{"你好呀", "客户端"})
}
