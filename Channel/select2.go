package main

import (
	"time"
	"fmt"
)

/*
Tuỳ thuộc channel nhận được dữ liệu trước thì
 */
func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string, 10)
	output2 := make(chan string, 10)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	//select sẽ block cho đến khi một trong các case nhận được dữ liệu từ channel
	fmt.Println("select finished blocking")
}
