package converter

import (
	"sarif-converter/sarifreport/report"
	"sarif-converter/sast"
)

type sastConverter struct {
}

var Sast = sastConverter{}

func (c sastConverter) Type() string {
	return "sast"
}

func (c sastConverter) Convert(input []byte) ([]byte, error) {
	sarifReport, err := report.FromBytes(input)
	if err != nil {
		return nil, err
	}

	r, err := sast.ConvertFrom(sarifReport)
	if err != nil {
		return nil, err
	}

	return r.Json()
}
