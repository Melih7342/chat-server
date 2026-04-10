package main

import (
	"github.com/Melih7342/chat-server/structs"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
}
