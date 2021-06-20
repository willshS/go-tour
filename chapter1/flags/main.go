package main

import (
	"flag"
	"fmt"
)

type myFlag struct {
	a int
	b string
}

func (f *myFlag) Set(value string) error {
	f.a = 1
	f.b = value
	return nil
}

func (f *myFlag) String() string {
	return fmt.Sprintf("myFlag a(%d) b(%s)", f.a, f.b)
}

func main() {
	var f myFlag
	flag.Var(&f, "myflags", "mmmmhelp")
	flag.Parse()
	fmt.Println(f)
}
