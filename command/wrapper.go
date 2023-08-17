package command

import (
	"sarif-converter/file"
	fake "sarif-converter/testing/file"
)

type wrapper struct {
	command Command
	io      file.IO
}

func (w wrapper) convert(args []string) {
	err := w.command.Convert(args)
	if err != nil {
		panic(err)
	}
}

func (w wrapper) readText(name string) string {
	b, err := w.io.Read(name)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func newWrapper() wrapper {
	io := fake.NewFakeIO()
	return wrapper{
		command: NewCommand(io, "", ""),
		io:      io,
	}
}
