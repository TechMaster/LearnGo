package main

import "fmt"

type Dog struct {
	Age interface{}
}

func main() {
	dog := Dog{}
	dog.Age = "3"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = 3
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = "not really an age"
	fmt.Printf("%#v %T", dog.Age, dog.Age)
}