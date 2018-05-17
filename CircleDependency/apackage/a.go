package apackage

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependency/bpackage"
)

type A struct {
	title string
	bObject *bpackage.B
}

func New (title string, b *bpackage.B) *A {
	return &A{
		title:   title,
		bObject: b,
	}
}

func (a *A) Bar() {
	fmt.Println(a.title)
	a.bObject.Foo()
}
