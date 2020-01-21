package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //We need a DB
)

//DB object
type Animal struct {
	ID    int
	Name  string
	AddDt string
}

//Get an aminal
func Get(id int) string {
	fmt.Println("Dal for animal")
	psqlConn := "host=localhost port=5432 user=postgres password=devpwd dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		fmt.Println("Could not connect to DB")
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not Ping DB")
		fmt.Println(err)
		panic(err)
	}

	var anAnimal string
	qerr := db.QueryRow(`SELECT name FROM animal WHERE id =$1`, id).Scan(&anAnimal)

	if qerr != nil {
		fmt.Println("query for animal failed")
		fmt.Println(err)
		panic(err)
	}

	return anAnimal
}
