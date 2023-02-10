package main

import (
	"codequality-converter/file"
	"fmt"
	"os"
)

var version = "unknown"
var revision = "unknown"

func main() {
	c := newConverterWith(file.NewIO())

	err := c.convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
