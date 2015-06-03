package fixture

import (
	"fmt"
	"os"
)
/*
four
line block comment
*/

/* stupid block *//* comment
spanning
*//* many lines in a weird way*/

//first single line comment
//send single line comment

type Foo struct {
}

type Bar struct {
	MyProperty int
}

func (b *Bar) ExportedMethod() {

}

func (b *Bar) unexportedMethod() {

}

type Baz interface {

}

func ExportedFunction() {
	os.Args
}

func unexportedFunction(){
	fmt.Print("")
}

const (
	FOO
	BAR
)

