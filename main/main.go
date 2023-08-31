package main

import (
	"fmt"
	"os"
	"sarif-converter/command"
	"sarif-converter/file"
	"sarif-converter/meta"
)

var version = "unknown"
var revision = "unknown"

func main() {
	c := command.NewCommand(file.NewIO(), meta.NewMetadata(version, revision))

	err := c.Convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
