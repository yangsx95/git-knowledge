package v1

import (
	"fmt"
	"git-knowledge/api/v1/vo"
)

// RepositoryApi 仓库操作API
type RepositoryApi interface {
}

type repositoryApiImpl struct {
}

func NewRepositoryApi() RepositoryApi {
	return &repositoryApiImpl{}
}

func (r *repositoryApiImpl) SayHello(ws vo.WebsocketSender, content string) error {
	fmt.Println("收到客户端消息：" + content)
	return ws.Send("你好呀，客户端")
}
