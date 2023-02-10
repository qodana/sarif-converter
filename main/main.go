package main

import (
	"codequality-converter/command"
	"codequality-converter/file"
	"fmt"
	"os"
)

func main() {
	c := command.NewCommand(file.NewIO())

	err := c.Convert(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
