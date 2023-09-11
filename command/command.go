package command

import (
	"errors"
	"fmt"
	"sarif-converter/converter"
	"sarif-converter/file"
	"sarif-converter/filter"
	"sarif-converter/main/argument"
	"sarif-converter/meta"
)

type Command struct {
	io       file.IO
	metadata *meta.Metadata
}

func (c Command) Convert(args []string) error {
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

func (c Command) parse(args []string) (*argument.Arguments, error) {
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

func (c Command) write(arguments *argument.Arguments, output []byte) error {
	return c.io.Write(arguments.Output(), output)
}

func (c Command) read(inputs file.Input) ([]byte, error) {
	return inputs.Read(c.io)
}

func (c Command) runFilter(input []byte, arguments *argument.Arguments) ([]byte, error) {
	return filter.AllSarifFilter(input, arguments)
}

func (c Command) runConvert(input []byte, arguments *argument.Arguments) ([]byte, error) {
	sarifConverter := converter.GetConverter(arguments.Type(), nil, c.metadata)
	return sarifConverter.Convert(input)

}

func (c Command) showVersion(arguments *argument.Arguments) {
	fmt.Println(arguments.Command() + " version " + c.version() + " (" + c.revision() + ")")
}

func (c Command) version() string {
	return c.metadata.Version
}

func (c Command) revision() string {
	return c.metadata.Revision
}

func NewCommand(io file.IO, metadata *meta.Metadata) Command {
	return Command{
		io:       io,
		metadata: metadata,
	}
}
