package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Animal")
	pathParams := mux.Vars(r)

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "This is an " + pathParams["id"]

	w.Write([]byte(message))
}
