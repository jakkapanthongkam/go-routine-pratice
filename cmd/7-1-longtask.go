package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ch chan int) {
	time.Sleep(6 * time.Second)
	ch <- 5
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan int)
	go longTask(ch)

	for {
		select {
		case msg := <-ch:
			fmt.Printf("Received: %d\n", msg)
		case <-ctx.Done():
			fmt.Println("Task cancelled/timed out:", ctx.Err())
			return
		default:
			fmt.Println("Task is working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
