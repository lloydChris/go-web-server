package main

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

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

	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
