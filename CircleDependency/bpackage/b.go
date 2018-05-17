package bpackage

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependency/apackage"
)

type B struct {
	name string
	aObject *apackage.A
}

func New (name string) *B {
	return &B {
		name: name,
	}
}

func (b *B) Foo() {
	fmt.Println(b.name)
}