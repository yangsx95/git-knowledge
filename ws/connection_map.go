package ws

import (
	"sync"
)

/*
记录所有用户长链接的Map集合
*/

type ConnectionMap struct {
	connections map[string]*Connection
	mutex       *sync.RWMutex
}

func NewConnectionMap() (cm *ConnectionMap) {
	cm = &ConnectionMap{connections: make(map[string]*Connection)}
	cm.mutex = &sync.RWMutex{}
	return
}

func (cm *ConnectionMap) Online(userid string, connection *Connection) {
	if connection == nil {
		return
	}
	cm.mutex.Lock()
	cm.connections[userid] = connection
	cm.mutex.Unlock()
}

func (cm *ConnectionMap) Offline(userid string) (connection *Connection) {
	cm.mutex.Lock()
	delete(cm.connections, userid)
	cm.mutex.Unlock()
	return
}

func (cm *ConnectionMap) Get(key string) (conn *Connection) {
	cm.mutex.RLock()
	conn = cm.connections[key]
	cm.mutex.RUnlock()
	return
}

func (cm *ConnectionMap) ForEach(handler func(connection *Connection)) {
	cm.mutex.RLock()
	for _, connection := range cm.connections {
		handler(connection)
	}
	cm.mutex.RUnlock()
}
