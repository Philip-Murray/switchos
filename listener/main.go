package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "karhoo-local")
	sub := client.Subscription("yeet")
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	custom_logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	for {
		time.Sleep(time.Second)
		pullMsgs(&sub)
	}
}

func pullMsgs(sub *Subscription) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	//ctx := context.Background()
	//client, err := pubsub.NewClient(ctx, "karhoo-local")
	//f err != nil {
	//return fmt.Errorf("pubsub.NewClient: %v", err)
	//}
	ctx := context.Background()
	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}
