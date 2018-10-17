package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "fresh-8-testing")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	topic := client.Topic("f8-testing-pubsub-watch")
	if topic == nil {
		log.Println("no topic")
		os.Exit(1)
	}

	sub, err := client.CreateSubscription(ctx, "my-sub-4", pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		if message.Attributes["eventType"] != "OBJECT_FINALIZE" {
			return
		}
		fileName := message.Attributes["objectId"]
		log.Printf("fileName: %s", fileName)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
