package main

import (
	"codequality-converter/converter"
	"fmt"
	"os"
)

var version = "unknown"
var revision = "unknown"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("codequality-converter version " + version + " (" + revision + ")")
		os.Exit(0)
	}

	input := tryRead()

	output := tryConvert(input)

	tryWrite(output)
}

func tryRead() []byte {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return input
}

func tryConvert(input []byte) []byte {
	output, err := converter.Convert(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return output
}

func tryWrite(output []byte) {
	err := os.WriteFile(os.Args[2], output, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
