package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alexsoup97/message-app/db"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	storage := db.InitalizeDb()
	defer storage.Db.Close()

	storage.CreateTables()
	router := chi.NewRouter()
	setupEndpoints(storage, router)

	fmt.Println("Server started...")
	http.ListenAndServe(":3000", router)

}
