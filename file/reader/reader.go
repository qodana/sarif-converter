package reader

type Reader interface {
	Read(name string) ([]byte, error)
}
