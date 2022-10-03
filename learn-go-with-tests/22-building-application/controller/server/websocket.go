package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// implements: io.Writer
type playerServerWS struct {
	ws *websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	ws, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Panicf("problem upgrading connection to Websockets %v\n", err)
	}

	return &playerServerWS{ws: ws}
}

// interface methods

func (this *playerServerWS) Write(p []byte) (n int, err error) {
	err = this.ws.WriteMessage(websocket.TextMessage, p)

	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// public methods

func (this *playerServerWS) WaitForMsg() string {
	_, msg, err := this.ws.ReadMessage()

	if err != nil {
		log.Panicf("error reading from Websocket %v\n", err)
	}

	return string(msg)
}
