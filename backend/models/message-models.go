package models

import "time"

type Message struct {
	Message         string `json:"message" validate:"required"`
	ConverstationId string `json:"conversationId" validate:"required"`
}

type Conversation struct {
	PermissionLevel int `json:"permission" validate:"required,number"`
}

type UserConversation struct {
	Name            string
	LastMessage     string
	LastMessageTime time.Time
}
