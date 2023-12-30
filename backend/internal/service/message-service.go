package service

import (
	"context"
	"database/sql"

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

func (message_serv MessageService) CreateConversation(
	convo *models.NewConversationRequest,
	ctx context.Context,
) (string, error) {

	convoSettings := models.ConversationSettings{
		ConversationName:  sql.NullString{String: convo.ConversationName, Valid: true},
		ConversationCount: sql.NullInt16{Int16: 2, Valid: true},
	}

	id, err := message_serv.database.MessageRepo.CreateConversationSettings(ctx, convoSettings)
	if err != nil {
		return "", err
	}

	var convoUsers [][]interface{}
	for _, user := range convo.Participants {
		convoUsers = append(convoUsers, []interface{}{id, user, 3})
	}

	convoUsers = append(convoUsers, []interface{}{id, ctx.Value("User"), 3})
	err = message_serv.database.MessageRepo.AddUsersToConversation(ctx, convoUsers)
	return id.String(), err
}

func (message_serv MessageService) CreateMessage(mssg *models.Message, ctx context.Context) error {
	err := message_serv.database.MessageRepo.SaveMessage(ctx, mssg)
	return err

}

func (message_serv MessageService) GetConversations(
	ctx context.Context,
) ([]models.UserConversation, error) {

	conversations, err := message_serv.database.MessageRepo.GetConversationsByUser(ctx)
	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (message_serv MessageService) GetChatHistory(ctx context.Context) {

	history, err := message_serv.database.MessageRepo.GetChatHistory()
}
