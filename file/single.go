package file

type Single struct {
	path string
}

func (s Single) Paths() []string {
	return []string{s.path}
}

func (s Single) IsEmpty() bool {
	return true
}

func (s Single) Read(reader Reader) ([]byte, error) {
	return reader.Read(s.path)
}

func newSingleFile(path string) Single {
	return Single{path: path}
}
