package main

import (
	"fmt"
	"time"
)
func queryComplexData(ch chan int) {
	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range arrayInt {
		time.Sleep(200 * time.Microsecond)
		ch <- val
	}
	close(ch)

}

func processComplexData(ch, quit chan int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		val, ok := <- ch
		if ok {
			fmt.Println(val)
		} else {
			quit <- 13
		}
	}


}
func main() {
	ch := make(chan int)
	quit := make(chan int)
	go queryComplexData(ch)
	go processComplexData(ch, quit)
	select {
		case <-quit:
			fmt.Println("It is done")
			return
			/* thử thêm đoạn code này vào,
			ta sẽ thấy main app không bị block để chờ tín hiệu quit luôn
			mà in ra Do something rồi thoát luôn
			default:
			fmt.Println("Do something")
			*/
	}

}
