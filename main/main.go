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

	input := tryRead(os.Args[1])

	output := tryConvert(input)

	tryWrite(output, os.Args[2])
}

func tryRead(input string) []byte {
	bytes, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}

func tryConvert(input []byte) []byte {
	output, err := converter.Convert(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return output
}

func tryWrite(output []byte, outputFile string) {
	err := os.WriteFile(outputFile, output, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
