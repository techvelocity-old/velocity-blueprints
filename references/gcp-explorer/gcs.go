package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"time"
)

func testGCS() {
	var vf bytes.Buffer
	if err := listFiles(&vf, mustGetenv("BUCKET_NAME")); err != nil {
		panic(err)
	}

	fmt.Println(vf.String())
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
