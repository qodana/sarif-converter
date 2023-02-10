package file

import (
	"errors"
)

type Empty struct {
}

func (n Empty) Paths() []string {
	return []string{}
}

func (n Empty) Read(Reader) ([]byte, error) {
	return nil, errors.New("null")
}

func (n Empty) IsEmpty() bool {
	return true
}

func newEmpty() Empty {
	return Empty{}
}
