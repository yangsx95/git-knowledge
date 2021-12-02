package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

/*
Connection

封装websocket连接对象，提供两个方法：
- ReadMessage 方法，负责将消息从 in channel中读取
- SendMessage 方法，负责将消息投递到 out channel

内部原理：
- 启动读协程，循环读取Websocket，将消息投递到 in channel
- 启动写协程，循环读取out channel，并将消息传递给websocket发送
*/
type Connection struct {
	wsConn  *websocket.Conn
	inChan  chan []byte
	outChan chan []byte
	// 当 readLoop 或者 writeLoop 其中一方出现问题导致websocket连接关闭时，
	// 另一方仍然处于从chan中读取/写入的状态中，协程将会阻塞在读取chan的步骤中；
	// 故设置关闭标记chan，当closeChan被关闭时，代表websocket连接也被关闭
	closeChan chan byte
	once      sync.Once
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte),
	}
	// 启动读写协程
	go conn.readLoop()
	go conn.writeLoop()
	return
}

// ReadMessage API
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) ReadJSON(target interface{}) (err error) {
	var message []byte
	if message, err = conn.ReadMessage(); err != nil {
		return
	}
	return json.Unmarshal(message, target)
}

// SendMessage API
func (conn *Connection) SendMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) SendJSON(data interface{}) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return
	}
	return conn.SendMessage(bytes)
}

// Close API
func (conn *Connection) Close() {
	// websocket.Conn 对象是线程安全的
	_ = conn.wsConn.Close()
	// websocket连接关闭，将closeChan标记也置于关闭状态
	conn.once.Do(func() {
		close(conn.closeChan)
	})
}

// 内部实现

// 循环读取Websocket，将消息投递到 in channel
func (conn *Connection) readLoop() {
	defer conn.Close()
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			return
		}
		select {
		case conn.inChan <- data: // 读到消息，放入inChain，如果读不到消息，就读取closeChan
		case <-conn.closeChan: // 如果读取不到消息代表websocket连接处于正常状态并进入下次循环，读取到消息说明websocket已经关闭
			// 备注：当从一个被关闭的通道中拿取数据时，如果通道中没有数据，该操作将不会再阻塞，而是返回一个false的状态位
			return
		}
	}
}

// 循环读取 out Chan 将信息传递给websocket发送
func (conn *Connection) writeLoop() {
	defer conn.Close()
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			return
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			return
		}
	}
}
