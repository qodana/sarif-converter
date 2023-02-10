package file

import (
	"codequality-converter/file/reader"
	"errors"
)

type Empty struct {
}

func (n Empty) Paths() []string {
	return []string{}
}

func (n Empty) Read(reader.Reader) ([]byte, error) {
	return nil, errors.New("null")
}

func (n Empty) IsEmpty() bool {
	return true
}

func newEmpty() Empty {
	return Empty{}
}
