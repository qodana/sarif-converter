package file

import "errors"

type Empty struct {
}

func (n Empty) Paths() []string {
	return []string{}
}

func (n Empty) Read() ([]byte, error) {
	return nil, errors.New("null")
}

func (n Empty) IsEmpty() bool {
	return true
}

func NewEmpty() Empty {
	return Empty{}
}
