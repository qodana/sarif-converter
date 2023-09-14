package sast

import (
	"sarif-converter/meta"
	"sarif-converter/now"
	"sarif-converter/sarifreport/report"
	"sarif-converter/sast"
)

type SastConverter struct {
	time     now.TimeProvider
	metadata meta.Metadata
}

var Sast = "sast"

func (c SastConverter) Convert(input []byte) ([]byte, error) {
	sarifReport, err := report.FromBytes(input)
	if err != nil {
		return nil, err
	}

	r, err := sast.ConvertFrom(sarifReport, c.time, c.metadata)
	if err != nil {
		return nil, err
	}

	return r.Json()
}

func NewSastConverter(time now.TimeProvider, metadata meta.Metadata) SastConverter {
	return SastConverter{
		time:     time,
		metadata: metadata,
	}
}
