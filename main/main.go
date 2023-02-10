package main

import (
	"codequality-converter/converter"
	"codequality-converter/file"
	"codequality-converter/file/reader"
	"codequality-converter/filter"
	"codequality-converter/main/argument"
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
		showVersion(arguments)
		os.Exit(0)
	}

	input := tryRead(arguments.Inputs())
	input = tryFilter(input, arguments)

	output := tryConvert(input, arguments)

	tryWrite(output, arguments.Output())
}

func showVersion(arguments *argument.Arguments) {
	fmt.Println(arguments.Command() + " version " + version + " (" + revision + ")")
}

func tryRead(input file.Input) []byte {
	bytes, err := input.Read(reader.NewReader())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}

func tryFilter(input []byte, arguments *argument.Arguments) []byte {
	converted, err := filter.AllSarifFilter(input, arguments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return converted
}

func tryConvert(input []byte, arguments *argument.Arguments) []byte {
	c := converter.GetConverter(arguments.Type())
	output, err := c.Convert(input)
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
