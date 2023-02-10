package file

type Reader interface {
	Read(name string) ([]byte, error)
}
