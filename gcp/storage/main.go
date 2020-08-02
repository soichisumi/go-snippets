package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

var (
	bucketName = ""
	ctx = context.Background()
)

func main() {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	it := client.Bucket(bucketName).Objects(ctx, &storage.Query{
		Delimiter: "/", // Delimiter returns results in a directory-like fashion.
						// Results will contain only objects whose names, aside from the prefix, do not contain delimiter.
		Prefix:    "2020-01-14/", // prefix of object path. ex: 2020-01-14/
		Versions:  false,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("err: %+v", err)
			return
		}
		fmt.Printf("obj: %s\n", attrs.Prefix)
	}
	log.Printf("done")
}
