package network

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ListenAndServe() {

	http.HandleFunc("/ws", upgradeWs)
	panic(http.ListenAndServe(":3001", nil))

}

func upgradeWs(w http.ResponseWriter, r *http.Request) {

	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	client := NewWebsocketClient(conn)

	go BeginReading(client)
	go BeginWriting(client)

}
