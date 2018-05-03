package main

import "fmt"
import "sync"

var wg sync.WaitGroup // 1

func routine() {
	//defer wg.Done() // 3 Chuyển wg.Done() xuống dưới cùng cũng được
	fmt.Println("routine finished")

}

func main() {
	wg.Add(1) // 2
	go routine() // *
	wg.Wait() // 4
	fmt.Println("main finished")
}