package main

import (
	"context"
	"fmt"
	"time"
)

func longTask1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled/timed out:", ctx.Err())
			return
		default:
			fmt.Printf("Task running\n")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	go longTask1(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine finished.")
}
