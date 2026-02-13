package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Sandesh30-cloud/go-user-api/internal/handler"
	"github.com/Sandesh30-cloud/go-user-api/internal/service"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
