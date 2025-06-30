package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting...\n", id)
	time.Sleep(time.Duration(id) * 1000 * time.Millisecond)
	fmt.Printf("Worker %d: Finished.\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Done\n")
}
