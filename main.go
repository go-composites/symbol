package main

import (
	"fmt"

	Symbol "github.com/go-composites/symbol/src"
)

func main() {
	foo := Symbol.New("foo")
	alsoFoo := Symbol.New("foo")
	bar := Symbol.New("bar")

	fmt.Println("Symbol.New(\"foo\").Inspect():           ", foo.Inspect())
	fmt.Println("New(\"foo\") == New(\"foo\") (interned):   ", foo == alsoFoo)
	fmt.Println("foo.Equal(alsoFoo):                     ", foo.Equal(alsoFoo))
	fmt.Println("foo.Equal(bar):                         ", foo.Equal(bar))
	fmt.Println("Symbol.Null().IsNull():                 ", Symbol.Null().IsNull())
}
