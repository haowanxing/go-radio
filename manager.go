package go_radio

// 客户端管理
type clientManager struct {
	clients    map[*client]bool
	broadcast  chan Message
	register   chan *client
	unregister chan *client
}

var manager = clientManager{
	broadcast:  make(chan Message),
	register:   make(chan *client),
	unregister: make(chan *client),
	clients:    make(map[*client]bool),
}

func startManager() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("manager.start defer:%+v", err)
		}
	}()
	for {
		select {
		case c := <-manager.register:
			manager.clients[c] = true
			log.Debugf("[Manager]（%s）加入服务器", c.id)
		case c := <-manager.unregister:
			if _, ok := manager.clients[c]; ok {
				delete(manager.clients, c)
				_ = c.Close()
			}
		case msg := <-manager.broadcast:
			log.Debugf("[Manager]广播信息：%v", msg)
			for conn := range manager.clients {
				conn.Send(msg)
			}
		}
	}
}
