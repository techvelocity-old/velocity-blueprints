package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func gets3Objects() *s3.ListObjectsOutput {
	svc := s3.New(awsSession)
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
		return nil
	}

	return result
}

func testS3() {
	log.Printf("S3 ListObjects: %v", gets3Objects())
}
