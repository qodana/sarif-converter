package reader

import "os"

type defaultReader struct {
}

func (d defaultReader) Read(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func NewReader() Reader {
	return defaultReader{}
}
