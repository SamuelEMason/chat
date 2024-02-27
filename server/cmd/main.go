package main

import (
	"chat/db"
	"chat/api"
	"chat/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// connection to db
	datab, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()
	router.HandleFunc("/users", api.GetUsers(datab.GetDB())).Methods("GET")
	router.HandleFunc("/users/{id}", api.GetUser(datab.GetDB())).Methods("GET")
	router.HandleFunc("/users", api.CreateUser(datab.GetDB())).Methods("POST")
	router.HandleFunc("/users/{id}", api.UpdateUser(datab.GetDB())).Methods("PUT")
	router.HandleFunc("/users/{id}", api.DeleteUser(datab.GetDB())).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", middleware.JsonContentTypeMiddleware(router)))
}