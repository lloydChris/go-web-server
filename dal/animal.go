package dal

import (
	"fmt"

	_ "github.com/lib/pq" //We need a DB
)

//Animal is the DB object
type Animal struct {
	ID    int
	Name  string
	AddDt string
}

//GetAnimal an animal
func (db *DB) GetAnimal(id int) (Animal, error) {
	var anAnimal Animal

	row := db.QueryRow(`SELECT id, name, add_dt FROM animal WHERE id =$1`, id)
	err := row.Scan(&anAnimal.ID, &anAnimal.Name, &anAnimal.AddDt)

	fmt.Printf("Retrived: %+v\n", anAnimal)
	if err != nil {
		fmt.Println("query for animal failed")
		fmt.Println(err)
	}

	return anAnimal, err
}
