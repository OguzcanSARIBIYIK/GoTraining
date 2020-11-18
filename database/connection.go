package database

import "database/sql"

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gogo")

	if err != nil {
		panic(err.Error())
	}

	return db
}
