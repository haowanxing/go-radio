package main

import (
	"github.com/gorilla/websocket"
	radio "go-radio"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/echo", wsHandle)
	radio.Run()
	log.Print(http.ListenAndServe(":1234", nil))
}

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
