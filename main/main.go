package main

import (
	"fmt"
	"os"
	"sarif-converter/command"
	"sarif-converter/file"
	"sarif-converter/meta"
)

func main() {
	c := command.NewCommand(file.NewIO(), meta.Version(), meta.Revision())

	err := c.Convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
