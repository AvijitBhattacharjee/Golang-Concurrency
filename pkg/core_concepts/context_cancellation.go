package coreconcepts

import (
	"context"
	"fmt"
	"time"
)


// How do you stop a goroutine? 
// How do you shutdown workers gracefully?

func worker_context(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("worker stopped")
			return

		default:
			fmt.Println("working...")
			time.Sleep(time.Second)
		}
	}
}

func ContextCancellation() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker_context(ctx)

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(time.Second)
}