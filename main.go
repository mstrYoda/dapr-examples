package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.PublishEvent(context.Background(), "rabbitmq-pubsub", "test-topic", []byte("test"))
	if err != nil {
		fmt.Println(err)
	}
}
