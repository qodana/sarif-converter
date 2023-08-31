package metadata

import _ "embed"

//go:embed sarif.json
var sarif []byte

//go:embed sast.json
var sast string

type Matadata struct {
}

func (s Matadata) Sarif() []byte {
	return sarif
}

func (s Matadata) Sast() string {
	return sast
}
