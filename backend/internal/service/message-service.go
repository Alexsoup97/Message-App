package service

import "github.com/Alexsoup97/message-app/db"

type MessageService struct {
	database *db.Storage
}

func (message_serv MessageService) CreateConversation() {
	message_serv.database.
}
