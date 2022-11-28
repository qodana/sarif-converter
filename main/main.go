package main

import (
	"codequality-converter/converter"
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
		usage()
		os.Exit(0)
	}

	if arguments.RequireShowVersion() {
		showVersion()
		os.Exit(0)
	}

	input := tryRead(arguments.Input())

	output := tryConvert(input)

	tryWrite(output, arguments.Output())
}

func usage() {
	fmt.Println("usage: codequality-converter input.sarif output.json")
	fmt.Println("  -v, --version\tdisplay version information and ext")
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
