package argument

import (
	"codequality-converter/file"
	"github.com/jessevdk/go-flags"
	"net/url"
	"os"
	"path/filepath"
)

type Arguments struct {
	args    []string
	options Options
	parser  *flags.Parser
	files   files
}

type Options struct {
	Version bool    `short:"v" long:"version" description:"Show version."`
	Type    string  `short:"t" long:"type" description:"Output report type." default:"html" choice:"sast" choice:"codequality" choice:"html"`
	SrcRoot *string `short:"r" long:"src-root" description:"Source root path."`
}

func (a Arguments) IsValid() bool {
	if a.RequireShowVersion() {
		return true
	}

	return a.files.isValid()
}

func (a Arguments) RequireShowVersion() bool {
	return a.options.Version
}

func (a Arguments) Inputs() file.Input {
	return file.NewInput(a.files.inputs)
}

func (a Arguments) Output() string {
	return a.files.output
}

func (a Arguments) Type() string {
	return a.options.Type
}

func (a Arguments) SrcRoot(basepath string) *string {
	if a.options.SrcRoot == nil {
		return nil
	}

	u, err := url.Parse(resolve(basepath, *a.options.SrcRoot))
	u.Scheme = "file"
	if err != nil {
		panic(err)
	}

	s := u.String()

	return &s
}

func resolve(basepath string, srcRoot string) string {
	if filepath.IsAbs(filepath.FromSlash(srcRoot)) {
		return srcRoot
	}

	return filepath.ToSlash(filepath.Join(basepath, srcRoot))
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
	parser.Usage = "[OPTIONS] input1.sarif [input2.sarif...] output.json"

	restArgs, err := parser.ParseArgs(args)
	if err != nil {
		return nil, err
	}

	files := parseFileArguments(restArgs[1:])
	return &Arguments{args: restArgs, options: options, parser: parser, files: files}, nil
}
