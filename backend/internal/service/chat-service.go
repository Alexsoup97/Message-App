package service

import (
	"github.com/Alexsoup97/message-app/models"
	"github.com/gorilla/websocket"
)

type ChatClient struct {
	Hub  *ChatHandler
	Conn *websocket.Conn
	User string
}

type ChatHandler struct {
	SendMessage chan *models.Message
	Register    chan *ChatClient
	Unregister  chan *ChatClient
	Clients     map[string]*ChatClient
	LiveChats   map[string][]*ChatClient
}

func (c ChatHandler) init() *ChatHandler {
	return &ChatHandler{
		SendMessage: make(chan *models.Message),
		Register:    make(chan *ChatClient),
		Unregister:  make(chan *ChatClient),
		Clients:     make(map[string]*ChatClient),
		LiveChats:   make(map[string][]*ChatClient),
	}

}

func (client ChatClient) read() {

	defer func() {
		client.Hub.Unregister <- &client
		client.Conn.Close()
	}()

	for {
		message := new(models.Message)
		err := client.Conn.ReadJSON(message)

		if err != nil {
			break
		}

		client.Hub.SendMessage <- message
	}
}

func (client ChatClient) write() {
	defer func() {

	}()

}

func (c ChatHandler) run() {
	for {
		select {
		case client := <-c.Register:
			c.Clients[client.User] = client

		case client := <-c.Unregister:
			delete(c.Clients, client.User)

		case message := <-c.SendMessage:

			targets, ok := c.LiveChats[message.ConverstationId]

			if !ok {

			}

			for _, s := range targets {

			}

		}
	}

}
