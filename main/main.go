package main

import (
	"codequality-converter/file/reader"
	"codequality-converter/file/writer"
	"fmt"
	"os"
)

var version = "unknown"
var revision = "unknown"

func main() {
	c := newConverter(reader.NewReader(), writer.NewWriter())

	err := c.convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
