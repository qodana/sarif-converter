package main

import (
	"codequality-converter/command"
	"codequality-converter/file"
	"fmt"
	"os"
)

var version = "unknown"
var revision = "unknown"

func main() {
	fmt.Println(version)
	c := command.NewCommand(file.NewIO(), version, revision)

	err := c.Convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
