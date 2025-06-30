package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(10000 * time.Millisecond)
	fmt.Printf("Hello Go routine\n")
}

func main() {
	go sayHello()
	fmt.Printf("Print from main func\n")
	time.Sleep(200 * time.Millisecond)
}
