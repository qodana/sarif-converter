package main

import (
	"codequality-converter/converter"
	"os"
)

func main() {

	input, _ := os.ReadFile(os.Args[1])

	output := converter.Convert(input)

	os.WriteFile(os.Args[2], output, 0666)
}
