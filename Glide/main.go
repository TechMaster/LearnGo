package main

import (
	"fmt"

	"github.com/TechMaster/funpackage"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
)

func main() {
	guid := xid.New()
	fmt.Println(guid.String())

	id := ksuid.New()
	fmt.Println(id.String())

	funpackage.Foo("Kinh đấy!!!")
}
