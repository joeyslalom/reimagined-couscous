package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	fmt.Println("Hello, world.")

	ctx := context.Background()
	projectId := "slalom-2020-293920"
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	topicId := "reimagined-couscous"
	t := client.Topic(topicId)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("ahnyoung"),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("result.Get: %v", err)
	}
	log.Printf("Published a message id=%v", id)
}