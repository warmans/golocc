package fixture

import (
	"fmt"
	"os"
)

/*
four
line block comment
*/

/* stupid block */ /* comment
spanning
*/ /* many lines in a weird way*/

//first single line comment
//send single line comment

type Foo struct {
}

type Bar struct {
	MyProperty int
}

func (b *Bar) ExportedMethod() {
	if 1 == 1 {
		return
	}

	if 1 == 2 {
		return
	} else {
		return
	}
}

func (b *Bar) unexportedMethod() {
	switch "foo" {
	case "bar":
		return
	case "baz":
		return
	}
}

type Baz interface {
}

func ExportedFunction(s string, i int64) {
	os.Args
}

func unexportedFunction() {
	go fmt.Print("")
	go fmt.Print("")

}

const FOO = "foo"
const BAR = "bar"
