package main

import "fmt"

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("Chạy sau")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 1333
	}()
	fmt.Println("Chạy trước")
	fibonacci2(c, quit)
}
