package file

import "os"

type IO interface {
	Read(name string) ([]byte, error)
	Write(name string, data []byte) error
}

type defaultIO struct {
}

func (io defaultIO) Read(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (io defaultIO) Write(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}

func NewIO() IO {
	return defaultIO{}
}
