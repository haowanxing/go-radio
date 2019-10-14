package go_radio

import "github.com/gorilla/websocket"

//  websocket客户端
type client struct {
	id      string
	socket  *websocket.Conn
	send    chan Message
	receive Receiver
}
type Receiver func(message Message)

func (c *client) ChangeReceiveFunction(fn Receiver) {
	c.receive = fn
}
func (c *client) ChangeID(id string) {
	c.id = id
}
func (c *client) Send(message Message) {
	c.send <- message
}
func (c *client) Close() error {
	close(c.send)
	return c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			manager.unregister <- c
			log.Errorf("发送消息给（%s）失败： ", c.id, err)
			break
		}
	}
	log.Infof("客户端（%s）关闭", c.id)
}

func (c *client) read() {
	for {
		var msg = Message{}
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Errorf("读取客户端消息失败：%s", err)
			manager.unregister <- c
			break
		}
		log.Debugf("收到来自（%s）的消息：%+v", c.id, msg)
		if dealFn := c.receive; dealFn != nil {
			dealFn(msg)
		}
	}
}
