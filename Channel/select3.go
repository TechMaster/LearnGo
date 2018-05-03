package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	select {
		case a:= <-ch:
			fmt.Println(a) //Nhận được sẽ thoát luôn
		case b:= <-ch:
			fmt.Println(b)
	}
	fmt.Println("Exit")

}