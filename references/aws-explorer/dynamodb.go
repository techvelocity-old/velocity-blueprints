package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func testDynamodb(session *session.Session) {
	dynamodbSvc := dynamodb.New(session)

	dynamodbResult, err := dynamodbSvc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(mustGetenv("DYNAMODB_TABLE_NAME")),
	})
	if err != nil {
		log.Fatalf("dynamodbSvc.DescribeTable: %v", err)
	}

	log.Printf("DynamoDB DescribeTable: %v", dynamodbResult)
}
