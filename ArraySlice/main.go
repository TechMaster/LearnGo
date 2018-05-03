package main

import (
	"fmt"
)

type Blah struct {
	f byte
}

func main() {
	//Khởi tạo mảng số nguyên
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	for index, value:= range primes {
		fmt.Printf("%d - %d \n",index, value)
	}

	//Khởi tạo slice kiểu nguyên
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[:3])


	//Khởi tạo slice của các kiểu khác nhau
	mixed := []interface{}{"foo", 10, &Blah{'c'}}
	for _, x := range mixed {
		fmt.Println(x)
	}

	fmt.Println(mixed[:2])


}