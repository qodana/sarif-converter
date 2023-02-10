package file

import "errors"

type FakeWriter struct {
	files map[string][]byte
}

func (r FakeWriter) Write(name string, data []byte) error {
	r.files[name] = data
	return nil
}

func (r FakeWriter) Get(name string) ([]byte, error) {
	data, ok := r.files[name]
	if ok {
		return data, nil
	}
	return nil, errors.New("`" + name + "` is not found")

}

func NewFakeWriter() FakeWriter {
	return FakeWriter{
		files: map[string][]byte{},
	}
}
