package routes

import (
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func CreateMessageRouter(storage *db.Storage) chi.Router {
	messageRouter := chi.NewRouter()

	messageRouter.Use(middleware.AuthMiddleware(storage))
	messageRouter.Get("/", test)
	return messageRouter
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}
