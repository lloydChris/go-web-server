package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Animal")

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "This is an " + message

	w.Write([]byte(message))
}
