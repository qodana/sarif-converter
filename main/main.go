package main

import (
	"codequality-converter/converter"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output, _ := converter.Convert(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(os.Args[2], output, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
