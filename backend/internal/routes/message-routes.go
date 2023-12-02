package routes

import (
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

func CreateMessageRouter(storage *db.Storage) chi.Router {
	messageRouter := chi.NewRouter()

	messageRouter.Use(middleware.AuthMiddleware(storage))
	messageRouter.Get("/heartbeat", heartbeat)
	return messageRouter
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
