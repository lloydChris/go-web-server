package dal

import (
	"database/sql"
	"fmt"
)

// Datastore in the Dal Interface
type Datastore interface {
	GetAnimal(int) (Animal, error)
}

// DB wrapper
type DB struct {
	*sql.DB
}

// NewDB returns a new DB connection
func NewDB(connectionString string) (*DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Could not connect to DB")
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not Ping DB")
		fmt.Println(err)
		return nil, err
	}

	return &DB{db}, nil
}
