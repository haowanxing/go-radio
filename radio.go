package go_radio

import (
	"github.com/d2r2/go-logger"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

var (
	log = logger.NewPackageLogger("radio", logger.InfoLevel)
)

// 消息类型
type Message struct {
	Event interface{} `json:"event"`
	Data  interface{} `json:"data"`
}

func Run() {
	go startManager()
}

// NewClient将Ws封装成client并注册到manager
// receiveFn用于自定义处理客户端发送的数据
func NewClient(conn *websocket.Conn, receiveFn Receiver) *client {
	var c = &client{
		id:      xid.New().String(),
		socket:  conn,
		send:    make(chan Message, 1024),
		receive: receiveFn,
	}
	manager.register <- c
	go c.read()
	go c.write()
	return c
}
func Broadcast(msg Message) {
	manager.broadcast <- msg
}
