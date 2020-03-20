package main

import (
	"go-web-server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/animal/{id}", handlers.GetAnimal).Methods("GET")

	http.ListenAndServe(":8081", router)
}
