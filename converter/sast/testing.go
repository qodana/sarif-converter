package sast

import (
	"sarif-converter/meta"
	"sarif-converter/now"
)

func NewSastConverterForTest() SastConverter {
	provider := now.NewFakeTime(now.Parse("2023-08-31T15:00:42Z"))
	metadata := meta.NewMetadata("0.5.1", "a9323")
	return NewSastConverter(provider, metadata)
}
