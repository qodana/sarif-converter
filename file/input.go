package file

type Input interface {
	IsEmpty() bool
	Read(reader Reader) ([]byte, error)
	Paths() []string
}

func NewInput(paths []string) Input {
	l := len(paths)

	if l <= 0 {
		return newEmpty()
	}

	if l == 1 {
		return newSingleFile(paths[0])
	}

	return newMultipleFiles(paths)
}
