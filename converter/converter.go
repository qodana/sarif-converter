package converter

type Converter interface {
	Convert(input []byte) ([]byte, error)
}
