package main

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependency/apackage"
	"github.com/TechMaster/LearnGo/CircleDependency/bpackage"
)

func main() {
	fmt.Println("Break dependency")
	b:= bpackage.New("This is B")
	a:= apackage.New("This is A", b)
	b.Foo()
	a.Bar()

}
