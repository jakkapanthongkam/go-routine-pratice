package main

import (
	"fmt"
	"sync"
	"time"
)

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numberOfJobs = 500
	const worker = 10
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	var wg sync.WaitGroup
	for w := 1; w <= worker; w++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			workerPool(i, jobs, results)
		}(w)
	}

	// Send jobs
	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result:", r)
	}

}
