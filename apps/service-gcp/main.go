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
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	_ "github.com/go-sql-driver/mysql"
)

var (
	topic *pubsub.Topic

	// Messages received by this instance.
	messagesMu sync.Mutex
	messages   []string
)

func main() {
	// GCS demo
	ctx := context.Background()
	var vf bytes.Buffer
	if err := listFiles(&vf, mustGetenv("BUCKET_NAME")); err != nil {
		panic(err)
	}

	fmt.Println(vf.String())

	// MySQL demo
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/dbname", mustGetenv("DB_USER"), mustGetenv("DB_PASSWORD"), mustGetenv("DB_HOST")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// PubSub demo
	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := mustGetenv("PUBSUB_TOPIC")
	topic = client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %v doesn't exist", topicName)
	}

	http.HandleFunc("/pubsub/publish", publishHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func listFiles(w io.Writer, bucket string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	msg := &pubsub.Message{
		Data: []byte(r.FormValue("payload")),
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Could not publish message: %v", err), 500)
		return
	}

	fmt.Fprint(w, "Message published.")
}
