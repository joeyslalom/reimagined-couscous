package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	subId := "reimagined-couscous-sub"
	sub := client.Subscription(subId)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		time.Sleep(10 * time.Second)
		m.Ack()
	})
	if err != nil {
		log.Fatalf("sub.Receive: %v", err)

	}
	cancel()
	log.Println("fin")
}
