package db

import (
	"context"

	"github.com/Alexsoup97/message-app/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepo struct {
	db *pgxpool.Pool
}

func (repo MessageRepo) SaveMessage(ctx context.Context, mssg *models.Message) error {
	query := `INSERT INTO messages 
	(MessageId, ConversationId, UserId, Message, createdAt)
	values($1, $2, $3,$4, NOW())
	`
	uuid := uuid.New()
	_, err := repo.db.Exec(ctx, query, uuid, mssg.ConverstationId, ctx.Value("User"), mssg.Message)
	return err
}

func (repo MessageRepo) CreateConversation(ctx context.Context, convo *models.Conversation) error {
	query := `INSERT INTO Conversations 
	(ConversationId, UserId, PermissionLevel)
	values($1, $2, $3) 
	`
	uuid := uuid.New()
	_, err := repo.db.Exec(ctx, query, uuid, ctx.Value("User"), convo.PermissionLevel)
	return err
}

func (repo MessageRepo) GetConversationsByUser(ctx context.Context) ([]models.UserConversation, error) {

	query := `SELECT (ConversationId, ConversationName, lastMessage) FROM Conversation_Settings LEFT JOIN Conversations 
	WHERE userid=$1 AND Conversation_Settings.ConversationId = Conversation.ConversationId`

	results, err := repo.db.Query(ctx, query, ctx.Value("User"))

	if err != nil {
		return nil, err
	}

	var conversations []models.UserConversation
	for results.Next() {

		var conversation models.UserConversation
		results.Scan(&conversation.conversationId)
		conversations = append(conversations, conversation)
	}
	return conversations, nil
}

func (s Storage) GetMessages() {

}
