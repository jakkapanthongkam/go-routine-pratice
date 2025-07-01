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
	defer rwMutex.RUnlock() // ควร defer RUnlock เพื่อให้มั่นใจว่าถูกปลดล็อคเสมอ

	time.Sleep(50 * time.Millisecond) // **ย้าย sleep มาก่อนการอ่าน**
	fmt.Printf("Reader %d: Reading data %d\n", id, data)
}

func writer(id int, val int) {
	rwMutex.Lock()
	defer rwMutex.Unlock() // ควร defer Unlock เพื่อให้มั่นใจว่าถูกปลดล็อคเสมอ

	fmt.Printf("Writer %d: Writing data %d\n", id, val)
	data = val
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
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
