package main

import (
	"errors"
	"fmt"
	"time"
)

func doRiskyTask(resultCh chan string, errCh chan error) {
	time.Sleep(100 * time.Millisecond)
	if time.Now().Second()%2 == 0 { // Simulate error sometimes
		errCh <- errors.New("something went wrong in risky task")
		return
	}
	resultCh <- "Task completed successfully!"
}

func main() {
	resultCh := make(chan string)
	errCh := make(chan error)

	go doRiskyTask(resultCh, errCh)

	select {
	case res := <-resultCh:
		fmt.Println("Success:", res)
	case err := <-errCh:
		fmt.Println("Error:", err)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout!")
	}
}
