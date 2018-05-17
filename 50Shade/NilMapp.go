package main

import (
	"fmt"
	"unicode/utf8"
)


func nilMap() {
	//Error
	var m map[string]int
	//Fix
	//var m map[string]int = make(map[string])

	// Or m := make(map[string]int)

	m["one"] = 1 //error
}

func indexString() {
	x := "text"
	fmt.Println(x[0]) //print 116
	fmt.Printf("%T\n",x[0]) //prints uint8

	//Correct way
	fmt.Println(string("Hello"[1]))              // ASCII only
	fmt.Println(string([]rune("Hello, 世界")[1])) // UTF-8
	fmt.Println(string([]rune("Hello, 世界")[8])) // UTF-8
}

func countString() {
	data := "♥"
	//Count bytes
	fmt.Println(len(data)) //prints: 3

	//Count character
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	data2 := "ô"
	fmt.Println(len(data2))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data2)) //prints: 2
}
func main() {
	//indexString()
	countString()
}