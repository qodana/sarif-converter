package writer

import "os"

type DefaultWriter struct {
}

func (w DefaultWriter) Write(name string, data []byte) error {
	return os.WriteFile(name, data, 0666)
}

func NewWriter() DefaultWriter {
	return DefaultWriter{}
}
