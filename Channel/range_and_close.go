package main

import (
	"fmt"
)
//hàm fibonacci  2 tham số: kích thước buffer, và channel với buffer tương ứng
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x  //gán dữ liệu vào channel
		x, y = y, x+y
	}
	close(c) //Chỉ có sender mới được close channel, receiver không nên tự close channel.
	//v, ok := <-ch dùng biến ok để biết xem
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {  //dùng for range để đọc ra từ channel.
		fmt.Println(i)
	}
}