package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Configure() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSON(w http.ResponseWriter, status int, object any) error {
	w.Header().Add("Content-Type", "application/json") 
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(object)
}
