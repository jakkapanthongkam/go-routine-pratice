package main

import (
	"fmt"
	"sync"
	"time"
)

func processData(input int) int {
	time.Sleep(1000 * time.Millisecond)
	return input * 10
}

func fanOut(input []int, out chan<- int) {
	defer close(out)
	var wg sync.WaitGroup

	for _, input := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out <- processData(input)
			fmt.Printf("Sending result :%d\n", input)
		}()
	}

	wg.Wait()
}

func fanIn(in <-chan int, done chan<- struct{}) {
	for result := range in {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Recevied result :%d\n", result)
	}
	done <- struct{}{}
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	for num := range 0 {
		input = append(input, num)
	}
	intermediate := make(chan int, len(input))
	done := make(chan struct{})

	go fanOut(input, intermediate)
	go fanIn(intermediate, done)

	<-done

	fmt.Printf("All done \n")
}
