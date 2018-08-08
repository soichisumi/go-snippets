package main

import (
	"context"
	"google.golang.org/api/option"
	"log"
	"firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"fmt"
)

func addData(ctx context.Context, client *firestore.Client) error{
	_, _, err := client.Collection("AddHere").Add(ctx, map[string]interface{}{
		"id": "yoyo",
	})
	return err
}

func batchedWrite(ctx context.Context, client *firestore.Client, datas []map[string]interface{}) error{
	batch := client.Batch()
	colRef := client.Collection("AddHere")
	for _, v := range datas {

	}


}

func getData(ctx context.Context, client *firestore.Client) error {
	snap, err := client.Doc(`GetFromHere/testdata`).Get(ctx)
	if err != nil {
		return err
	}
	data := snap.Data()
	fmt.Println(data)
	return err
}

func main() {
	opt := option.WithCredentialsFile("/Users/soichisumi/.secrets/fir-test-2fc32-firebase-adminsdk-94540-0742eb53e2.json")
	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	client, err := app.Firestore(ctx)
	addData(ctx, client)
	getData(ctx, client)
	defer client.Close()
}
