package main

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
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
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		log.Fatalf("session.NewSession: %v", err)
	}

	// S3 Demo
	svc := s3.New(newSession)
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(mustGetenv("BUCKET_NAME")),
		MaxKeys: aws.Int64(2),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				log.Fatalf("Unhandled error: %v", err)
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatalf("Unhandled error: %v", err)
		}
		return
	}

	log.Printf("S3 ListObjects: %v", result)

	// SQS demo
	sqsSvc := sqs.New(newSession)

	sqsResult, err := sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(mustGetenv("QUEUE_NAME")),
	})
	if err != nil {
		log.Fatalf("sqsSvc.GetQueueUrl: %v", err)
	}

	log.Printf("Queue URL: %v\n", *sqsResult.QueueUrl)

	// PSQL demo
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
	select {}
}
