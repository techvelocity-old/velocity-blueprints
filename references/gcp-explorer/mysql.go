package main

import (
	"database/sql"
	"fmt"
)

func testMySQL() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/app_development", mustGetenv("DB_USER"), mustGetenv("DB_PASS"), mustGetenv("DB_HOST")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
