package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	testGCS()
	testMySQL()
	testPubSub()

	log.Println("All cloud resources tests passed!")

	select {}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
