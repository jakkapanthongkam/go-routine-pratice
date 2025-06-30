package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	fmt.Println("Buffered Producer: Sending 1")
	ch <- 1
	fmt.Println("Buffered Producer: Sending 2")
	ch <- 2
	fmt.Println("Buffered Producer: Sending 3")
	ch <- 3
	close((ch))
}

func main() {
	ch := make(chan int, 5)

	go producer(ch)
	fmt.Printf("main \n")
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Main received %d\n", <-ch)
	fmt.Printf("Main received %d\n", <-ch)
	fmt.Printf("Main received %d\n", <-ch)

	fmt.Printf("Main Done\n")
}

// Note
// ถ้า channel มากกว่า ก็เป็นการสร้าง buffer เปลืองเฉยๆ
// ถ้า น้อยกว่า จะโดนตัดสัญญาน
