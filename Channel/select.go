package main

import "fmt"
/*
See more at https://github.com/t9md/go-study/blob/master/hello/go-tour-select.go
 */
func fibonacci2(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			fmt.Printf("Assign %d to c\n", x)
			x, y = y, x+y

		case <-quit:  //
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
			fmt.Println("Read")
			fmt.Println(<-c)  //đọc ra từ channel c
		}
		quit <- 1333  //Gửi message vào channel quit
	}()
	fmt.Println("Chạy trước")
	fibonacci2(c, quit)
}
