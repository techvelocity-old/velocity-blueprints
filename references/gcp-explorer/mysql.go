package main

import (
	"database/sql"
	"fmt"
)

func testMySQL() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/dbname", mustGetenv("DB_USER"), mustGetenv("DB_PASSWORD"), mustGetenv("DB_HOST")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
