package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"net/http"
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

var (
	awsSession *session.Session
)

func main() {
	var err error
	awsSession, err = session.NewSession(&aws.Config{
		Region: aws.String(mustGetenv("AWS_REGION"))},
	)
	if err != nil {
		log.Fatalf("session.NewSession: %v", err)
	}

	testS3()
	testSQS()
	testDynamodb()
	testPSQL()

	log.Println("HTTP Server starting...")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

	select {}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello!\n")
	s3Objects := gets3Objects()
	s3ObjectsJSON, err := json.Marshal(s3Objects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "s3 objects: %v", string(s3ObjectsJSON))
}
