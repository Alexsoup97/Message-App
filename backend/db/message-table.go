package db

import (
	"context"
	"time"

	"github.com/Alexsoup97/message-app/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepo struct {
	db *pgxpool.Pool
}

func (repo MessageRepo) SaveMessage(ctx context.Context, mssg *models.Message) error {
	query := `INSERT INTO messages 
	(message_id, conversation_id, user_id, message, created_at)
	values($1, $2, $3,$4, NOW())
	`
	uuid := uuid.New()
	_, err := repo.db.Exec(ctx, query, uuid, mssg.ConverstationId, ctx.Value("User"), mssg.Message)
	return err
}

func (repo MessageRepo) AddUsersToConversation(
	ctx context.Context,
	usersToAdd [][]interface{},
) error {

	_, err := repo.db.CopyFrom(
		ctx,
		pgx.Identifier{"conversations"},
		[]string{"conversation_id", "user_id", "permission_level"},
		pgx.CopyFromRows(usersToAdd),
	)

	return err
}

func (repo MessageRepo) CreateConversationSettings(
	ctx context.Context,
	settings models.ConversationSettings,
) (uuid.UUID, error) {

	query := `INSERT INTO conversation_settings(conversation_id, conversation_name, conversation_type,created_at, last_message, last_message_time) values($1, $2, $3, $4,'created group chat', $4)`
	id := uuid.New()
	_, err := repo.db.Exec(ctx, query, id, settings.ConversationName, 3, time.Now())
	return id, err
}

func (repo MessageRepo) GetConversationsByUser(
	ctx context.Context,
) ([]models.UserConversation, error) {

	query := `SELECT convos.conversation_id, conversation_name, last_message, last_message_time FROM (SELECT * FROM conversations WHERE user_id=$1) as convos
	 RIGHT JOIN conversation_settings on	conversation_settings.conversation_id = convos.conversation_id WHERE convos.conversation_id IS NOT NULL`

	results, err := repo.db.Query(ctx, query, ctx.Value("User"))

	defer results.Close()
	if err != nil {
		return nil, err
	}

	var conversations []models.UserConversation
	for results.Next() {
		var convo models.UserConversation
		err := results.Scan(
			&convo.ConversationId,
			&convo.Name,
			&convo.LastMessage,
			&convo.LastMessageTime,
		)

		if err != nil {
			return nil, err
		}
		conversations = append(conversations, convo)
	}
	return conversations, nil
}

func (s Storage) GetMessages() {

}
