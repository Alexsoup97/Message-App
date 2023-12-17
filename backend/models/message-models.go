package models

import (
	"database/sql"
	"time"
)

type Message struct {
	Message         string `json:"message"        validate:"required"`
	ConverstationId string `json:"conversationId" validate:"required"`
}

type NewConversationRequest struct {
	Participants     []string `json:"participants"`
	ConversationName string   `json:"name"         validate:"required"`
}

type Conversation struct {
	PermissionLevel int
	User            string
	ConversationId  string
}

type UserConversation struct {
	Name            string
	LastMessage     string
	LastMessageTime time.Time
	ConversationId  string
}

type ConversationSettings struct {
	ConversationId    sql.NullString
	ConversationName  sql.NullString
	ConversationCount sql.NullInt16
	LastMessage       sql.NullString
	LastMessageTime   sql.NullTime
}
