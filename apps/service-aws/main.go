// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample pubsub demonstrates use of the cloud.google.com/go/pubsub package from App Engine flexible environment.
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
		panic(err)
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
				panic(aerr)
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			panic(err)
		}
		return
	}

	fmt.Println(result)

	// SQS demo
	sqsSvc := sqs.New(newSession)

	sqsResult, err := sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(mustGetenv("QUEUE_NAME")),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(sqsResult)

	// PSQL demo
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		mustGetenv("DB_HOST"), mustGetenv("DB_PORT"), mustGetenv("DB_USER"), mustGetenv("DB_PASS"), "db")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
