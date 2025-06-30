package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i <= 3; i++ {
		if i == 2 {
			time.Sleep(5 * time.Second)
		}
		ch <- i
		fmt.Printf("Producer send %d\n", i)
	}

	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for val := range ch {
		fmt.Printf("Received %d\n", val)
	}

	fmt.Print("Done !!")
}
