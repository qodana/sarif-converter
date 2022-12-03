package converter

type Converter interface {
	Type() string
	Convert(input []byte) ([]byte, error)
}
