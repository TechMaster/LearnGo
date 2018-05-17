package main

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependencySolved/apackage"
	"github.com/TechMaster/LearnGo/CircleDependencySolved/bpackage"
)

func main() {
	fmt.Println("Break dependency")
	b:= bpackage.New("This is B")
	a:= apackage.New("This is A", b)
	b.SetAObject(a)
	//b.Foo()
	a.Bar()
	//b.DoBar()

}
