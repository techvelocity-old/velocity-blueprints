package main

import (
	"database/sql"
	"fmt"
	"log"
)

func testPSQL() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		mustGetenv("DB_HOST"), mustGetenv("DB_PORT"), mustGetenv("DB_USER"), mustGetenv("DB_PASS"), "app_development")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("db.Ping: %v", err)
	}

	log.Println("DB connection successful!")
}
