package bpackage

import (
	"fmt"
	"github.com/TechMaster/LearnGo/CircleDependencySolved/share"
)

type B struct {
	name string
	aObject share.AInterface
}

func New (name string) *B {
	return &B {
		name: name,
	}
}

func (b *B) SetAObject(a share.AInterface) {
	b.aObject = a
}

func (b *B) Foo() {
	fmt.Println("func (b *B) Foo()", b.name)
}

func (b *B) DoBar() {
	b.aObject.Bar()
}

