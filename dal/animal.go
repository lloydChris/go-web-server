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
func Get(id int) Animal {
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

	var anAnimal Animal

	row := db.QueryRow(`SELECT id, name, add_dt FROM animal WHERE id =$1`, id)
	err = row.Scan(&anAnimal.ID, &anAnimal.Name, &anAnimal.AddDt)

	fmt.Printf("Retrived: %+v", anAnimal)
	if err != nil {
		fmt.Println("query for animal failed")
		fmt.Println(err)
		panic(err)
	}

	return anAnimal
}
