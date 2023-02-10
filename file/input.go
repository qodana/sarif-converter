package file

type Input interface {
	IsEmpty() bool
	Read() ([]byte, error)
	Paths() []string
}
