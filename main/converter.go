package main

import (
	"codequality-converter/converter"
	"codequality-converter/file"
	"codequality-converter/file/reader"
	"codequality-converter/file/writer"
	"codequality-converter/filter"
	"codequality-converter/main/argument"
	"errors"
	"fmt"
)

type Converter struct {
	reader reader.Reader
	writer writer.Writer
}

func (c Converter) convert(args []string) error {
	arguments, err := c.parse(args)
	if err != nil {
		return err
	}

	if arguments.RequireShowVersion() {
		c.showVersion(arguments)
		return nil
	}

	input, err := c.read(arguments.Inputs())
	if err != nil {
		return err
	}

	input, err = c.runFilter(input, arguments)
	if err != nil {
		return err
	}

	output, err := c.runConvert(input, arguments)
	if err != nil {
		return err
	}

	return c.write(arguments, output)
}

func (c Converter) parse(args []string) (*argument.Arguments, error) {
	arguments, err := argument.Parse(args)
	if err != nil {
		return nil, err
	}

	if !arguments.IsValid() {
		arguments.ShowUsage()
		return nil, errors.New("invalid command line arguments")
	}

	return arguments, nil
}

func (c Converter) write(arguments *argument.Arguments, output []byte) error {
	return c.writer.Write(arguments.Output(), output)
}

func (c Converter) read(inputs file.Input) ([]byte, error) {
	return inputs.Read(c.reader)
}

func (c Converter) runFilter(input []byte, arguments *argument.Arguments) ([]byte, error) {
	return filter.AllSarifFilter(input, arguments)
}

func (c Converter) runConvert(input []byte, arguments *argument.Arguments) ([]byte, error) {
	sarifConverter := converter.GetConverter(arguments.Type())
	return sarifConverter.Convert(input)

}

func (c Converter) showVersion(arguments *argument.Arguments) {
	fmt.Println(arguments.Command() + " version " + version + " (" + revision + ")")
}

func newConverter(reader reader.Reader, writer writer.Writer) Converter {
	return Converter{
		reader: reader,
		writer: writer,
	}
}
