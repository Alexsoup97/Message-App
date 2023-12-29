package models

type ErrorResponse struct {
	Message string `json:"error"`
}

type HeartBeatResponse struct {
	Username string `json:"user"`
}

type CreateConversationResponse struct {
	ConversationId string `json:"id"`
}
