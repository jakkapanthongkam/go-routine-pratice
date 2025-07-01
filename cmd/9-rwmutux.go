package main

import (
	"fmt"
	"sync"
	"time"
)

var data int
var rwMutex sync.RWMutex

func reader(id int) {
	rwMutex.RLock()
	fmt.Printf("Reader %d: Reading data %d\n", id, data)
	time.Sleep(10000 * time.Millisecond)
	rwMutex.RUnlock()
}

func writer(id int, val int) {
	rwMutex.Lock()
	fmt.Printf("Writer %d: Writing data %d\n", id, val)
	data = val
	time.Sleep(20000 * time.Millisecond)
	rwMutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			reader(i)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		writer(1, 100)
	}()

	wg.Wait()
	fmt.Println("All read/write operations finished.")
}
