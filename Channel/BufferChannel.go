package main

import "fmt"

func main() {
	/*
	Nếu buffer length quá thấp sẽ gặp phải lỗi này
	fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [chan send]:
	main.main()

	Sends to a buffered channel block only when the buffer is full.
	Receives block when the buffer is empty.
	 */
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}