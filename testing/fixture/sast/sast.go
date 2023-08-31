package sast

import (
	"sarif-converter/testing/fixture/sast/metadata"
)

type Sast struct {
	Metadata metadata.Matadata
}

func NewSast() Sast {
	return Sast{}
}
