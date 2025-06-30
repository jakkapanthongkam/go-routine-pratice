package main

import (
	"fmt"
	"time"
)

func doSomething(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go doSomething(ch1, "Message from Channel 1", 10000*time.Millisecond)
	go doSomething(ch2, "Message from Channel 2", 50*time.Millisecond)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("%s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("%s\n", msg2)
		}
	}

	fmt.Printf("Done\n")
}
