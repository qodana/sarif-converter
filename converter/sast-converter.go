package converter

import (
	report2 "sarif-converter/sarifreport/report"
	sast2 "sarif-converter/sast"
)

type sastConverter struct {
}

var Sast = sastConverter{}

func (c sastConverter) Type() string {
	return "sast"
}

func (c sastConverter) Convert(input []byte) ([]byte, error) {
	sarifReport, err := report2.FromBytes(input)
	if err != nil {
		return nil, err
	}

	return sast2.ConvertFrom(sarifReport)
}
