package apackage

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependencySolved/share"
)

type A struct {
	title string
	bObject share.BInterface
}

func New (title string, foo share.BInterface) *A {
	return &A{
		title:   title,
		bObject: foo,
	}
}

func (a *A) Bar() {
	fmt.Println("func (a *A) Bar()", a.title)
	a.bObject.Foo()
}
