package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/middleware"
	"github.com/Alexsoup97/message-app/models"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type IMessageService interface {
	CreateMessage(mssg *models.Message, reqContext context.Context) error
	CreateConversation(
		conversation *models.NewConversationRequest,
		reqContext context.Context,
	) (string, error)
	GetConversations(ctx context.Context) ([]models.UserConversation, error)
}

type MessageRouter struct {
	messageService IMessageService
}

func CreateMessageRouter(storage *db.Storage, message_service IMessageService) chi.Router {
	messageRouter := chi.NewRouter()

	router := &MessageRouter{
		messageService: message_service,
	}

	messageRouter.Use(middleware.AuthMiddleware(storage))
	messageRouter.Get("/heartbeat", heartbeat)
	messageRouter.Post("/send", router.createMessage)
	messageRouter.Post("/conversations", router.createConversation)
	messageRouter.Get("/conversations", router.getConversations)
	return messageRouter
}

func (router MessageRouter) createMessage(w http.ResponseWriter, r *http.Request) {

	req := new(models.Message)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Print(err)
	}

	err := validate.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = router.messageService.CreateMessage(req, r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (router MessageRouter) createConversation(w http.ResponseWriter, r *http.Request) {

	req := new(models.NewConversationRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Print(err)
	}

	err := validate.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convo_id, err := router.messageService.CreateConversation(req, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusCreated, models.CreateConversationResponse{ConversationId: convo_id})

}

func (router MessageRouter) getConversations(w http.ResponseWriter, r *http.Request) {

	convos, err := router.messageService.GetConversations(r.Context())

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, convos)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	WriteJSON(
		w,
		http.StatusOK,
		models.HeartBeatResponse{Username: r.Context().Value("User").(string)},
	)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
