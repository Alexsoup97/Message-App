package service

import (
	"context"

	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/models"
)

type MessageService struct {
	database *db.Storage
}

func CreateMessageService(db *db.Storage) *MessageService {
	return &MessageService{
		database: db,
	}
}
func (message_serv MessageService) CreateConversation(convo *models.Conversation, ctx context.Context) error {

	err := message_serv.database.MessageRepo.CreateConversation(ctx, convo)
	return err
}

func (message_serv MessageService) CreateMessage(mssg *models.Message, ctx context.Context) error {
	err := message_serv.database.MessageRepo.SaveMessage(ctx, mssg)
	return err

}
