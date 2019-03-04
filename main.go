package main

import (
	"database/sql"
	"go-web-server/handlers"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	psqlConn := "host=localhost port=5432 user=postgres password=devpwd dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/animal", handlers.GetAnimal).Methods("GET")

	http.ListenAndServe(":8081", router)
}
