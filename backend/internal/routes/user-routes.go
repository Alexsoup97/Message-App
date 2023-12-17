package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/Alexsoup97/message-app/internal/service"
	"github.com/Alexsoup97/message-app/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRouter struct {
	userService IUserService
}

type IUserService interface {
	Login(username string, passwor string) (string, error)
	CreateAccount(username string, password string) error
}

func CreateUserRouter(userService IUserService) chi.Router {

	userRouter := &UserRouter{
		userService: userService,
	}

	chiRouter := chi.NewRouter()
	chiRouter.Post("/", userRouter.login)
	chiRouter.Post("/create", userRouter.createAccount)

	return chiRouter
}

func (router UserRouter) createAccount(w http.ResponseWriter, r *http.Request) {
	user, pass := r.FormValue("username"), r.FormValue("password")
	if user == "" || pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := router.userService.CreateAccount(user, pass)

	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		WriteJSON(
			w,
			http.StatusBadRequest,
			models.ErrorResponse{Message: "User already exist. Please enter a new username"},
		)
		return
	}

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (router UserRouter) login(w http.ResponseWriter, r *http.Request) {
	user, pass := r.FormValue("username"), r.FormValue("password")

	token, err := router.userService.Login(user, pass)

	if err != nil {
		switch err {
		case service.IncorrectPassword, service.UserNotFound:
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		}
		return
	}

	cookie := &http.Cookie{
		Name:     "APIAuth",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (router UserRouter) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
