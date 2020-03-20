package main

import (
	"fmt"
	"go-web-server/dal"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Env struct {
	db dal.Datastore
}

func main() {
	db, err := dal.NewDB("host=localhost port=5432 user=postgres password=devpwd dbname=postgres sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	router := mux.NewRouter()

	router.HandleFunc("/animal/{id}", env.GetAnimal).Methods("GET")

	http.ListenAndServe(":8081", router)
}

//GetAnimal an Animal
func (env *Env) GetAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Animal")
	pathParams := mux.Vars(r)

	animalID, _ := strconv.Atoi(pathParams["id"])
	fmt.Println(animalID)
	anAnimal, err := env.db.GetAnimal(animalID)
	if err != nil {
		panic(err)
	}

	message := "This is an " + anAnimal.Name

	w.Write([]byte(message))
}
