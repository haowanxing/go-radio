# Websocket广播通信库（轮子）

基于此项目，可以很快实现诸如聊天室、广播站等冇趣的玩意儿！

使用方法：
--------

```go
// 设置允许跨域*
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main(){
	http.HandleFunc("/echo", wsHandle)
	radio.Run()
	log.Print(http.ListenAndServe(":1234", nil))
}

// 定义websocket通信方法
func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	radio.NewClient(conn, func(msg radio.Message) {
		radio.Broadcast(msg)
	})
}
```

