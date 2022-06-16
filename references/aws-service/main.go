package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func main() {
	newSession, err := session.NewSession(&aws.Config{
		Region: aws.String(mustGetenv("AWS_REGION"))},
	)
	if err != nil {
		log.Fatalf("session.NewSession: %v", err)
	}

	tests3(newSession)
	testSQS(newSession)
	testDynamodb(newSession)
	testPSQL()

	select {}
}
