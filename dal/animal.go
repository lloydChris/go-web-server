package dal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Animal struct {
	id     int
	name   string
	add_dt string
}

func Get(id int) {
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

	anAnimal = make(Animal)
	row := db.QueryRow(`SELECT * FROM animal WHERE id =$1`, id)

	return sql.row
}
