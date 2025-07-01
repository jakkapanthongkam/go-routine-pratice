package main

import (
	"fmt"
	"sync"
)

var counter int
var mux sync.Mutex

func increment() {
	mux.Lock()
	defer mux.Unlock()

	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Done at %d\n", counter)
}
