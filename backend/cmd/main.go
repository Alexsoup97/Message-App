package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	storage := db.InitalizeDb()
	defer storage.Db.Close()

	storage.CreateUserTables()
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))
	setupEndpoints(storage, router)
	frontendConfigure(router)

	fmt.Println("Server started...")
	http.ListenAndServe(":3000", router)

}
