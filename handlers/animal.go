package handlers

import (
	"fmt"
	"go-web-server/dal"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAnimal an Animal
func GetAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Animal...")
	pathParams := mux.Vars(r)

	animalID, _ := strconv.Atoi(pathParams["id"])
	fmt.Println(animalID)
	anAnimal := dal.Get(animalID)

	fmt.Println(anAnimal)
	message := "This is an " + anAnimal.Name

	w.Write([]byte(message))
}
