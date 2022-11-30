package argument

import (
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
)

type Arguments struct {
	args    []string
	options Options
	parser  *flags.Parser
}

type Options struct {
	Version bool   `short:"v" long:"version" description:"Show version."`
	Type    string `short:"t" long:"type" description:"Output report type." default:"codequality" choice:"sast" choice:"codequality" choice:"html"`
}

func (a Arguments) IsValid() bool {
	if a.RequireShowVersion() {
		return true
	}

	if a.Input() != "" && a.Output() != "" {
		return true
	}

	return false
}

func (a Arguments) RequireShowVersion() bool {
	return a.options.Version
}

func (a Arguments) Input() string {
	if len(a.args) > 1 {
		return a.args[1]
	}
	return ""
}

func (a Arguments) Output() string {
	if len(a.args) > 2 {
		return a.args[2]
	}
	return ""
}

func (a Arguments) Type() string {
	return a.options.Type
}

func (a Arguments) ShowUsage() {
	a.parser.WriteHelp(os.Stdout)
}

func (a Arguments) Command() string {
	return a.parser.Name
}

func Parse(args []string) (*Arguments, error) {
	options := Options{}
	parser := flags.NewParser(&options, flags.Default)
	parser.Name = filepath.Base(args[0])
	parser.Usage = "[OPTIONS] input.sarif output.json"

	restArgs, err := parser.ParseArgs(args)
	if err != nil {
		return nil, err
	}
	return &Arguments{args: restArgs, options: options, parser: parser}, nil
}
