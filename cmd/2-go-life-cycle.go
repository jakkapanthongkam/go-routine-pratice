package main

import (
	"fmt"
	"time"
)

func doWork() {
	for i := 0; i <= 100; i++ {
		// child process will die after main bye.
		// time.Sleep(5000 * time.Microsecond)
		time.Sleep(50 * time.Microsecond)
		fmt.Printf("Working... %d\n", i)
	}
}

func main() {
	go doWork()
	fmt.Printf("Main goroutine continues...\n")

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Main say good bye\n")
}
