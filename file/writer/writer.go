package writer

type Writer interface {
	Write(name string, data []byte) error
}
