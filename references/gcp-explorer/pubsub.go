package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func testPubSub() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := mustGetenv("PUBSUB_TOPIC")
	topic := client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %v doesn't exist", topicName)
	}
}
