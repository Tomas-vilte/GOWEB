package main

import (
	"apirest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Rutas
	mux := mux.NewRouter()

	// Endpoints
	mux.HandleFunc("/api/user/", controllers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", controllers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
