package routes

import (
	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

func CreateMessageRouter(storage *db.Storage) chi.Router {
	messageRouter := chi.NewRouter()

	messageRouter.Use(middleware.AuthMiddleware(storage))
	return messageRouter
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
