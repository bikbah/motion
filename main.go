package main

import (
	"fmt"
	"os"
)

var src = `package main

var bar = func() {}

func foo() error {
	_ = func() {
		// -------
	}
	return nil
}
`

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	funcs, err := NewFuncs([]byte(src))
	if err != nil {
		return err
	}

	for _, fn := range funcs {
		fmt.Println(fn)
	}
	return nil
}