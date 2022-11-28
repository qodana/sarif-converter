package main

import (
	"codequality-converter/converter"
	"codequality-converter/main/argument"
	"errors"
	"fmt"
	"os"
)

var version = "unknown"
var revision = "unknown"

func main() {
	arguments, err := argument.Parse(os.Args)
	if err != nil {
		panic(err)
	}

	if !arguments.IsValid() {
		arguments.ShowUsage()
		os.Exit(1)
	}

	if arguments.RequireShowVersion() {
		showVersion()
		os.Exit(0)
	}

	input := tryRead(arguments.Input())

	output := tryConvert(input, arguments)

	tryWrite(output, arguments.Output())
}

func showVersion() {
	fmt.Println("codequality-converter version " + version + " (" + revision + ")")
}

func tryRead(input string) []byte {
	bytes, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}

func tryConvert(input []byte, arguments *argument.Arguments) []byte {
	if arguments.Type() == "codequality" {
		output, err := converter.Convert(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return output
	}
	if arguments.Type() == "sast" {
		output, err := converter.ConvertToSast(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return output
	}

	panic(errors.New("invalid report type"))
}

func tryWrite(output []byte, outputFile string) {
	err := os.WriteFile(outputFile, output, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
