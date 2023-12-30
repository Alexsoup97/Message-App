package service

import "github.com/gorilla/websocket"

type ChatClient struct {
	Hub  *ChatHandler
	Conn *websocket.Conn
  User string
}

type ChatHandler struct {
	SendMessage chan []byte
	Register    chan *ChatClient
	Unregister  chan *ChatClient
	Clients     map[string]*ChatClient
}

func (c ChatHandler) init() *ChatHandler {
	return &ChatHandler{
		SendMessage: make(chan []byte),
		Register:    make(chan *ChatClient),
		Unregister:  make(chan *ChatClient),
		Clients:     make(map[string]*ChatClient),
	}

}

func (c ChatHandler) run() {
	for {
		select {
		case client := <-c.Register:
      c.Clients[client.]


		case client := <-c.Unregister:

		case message := <-c.SendMessage:

		}
	}

}
