package converter

import (
	"sarif-converter/codequality"
	"sarif-converter/sarifreport/report"
)

type codeQualityConverter struct {
}

func (c codeQualityConverter) Type() string {
	return "codequality"
}

func (c codeQualityConverter) Convert(input []byte) ([]byte, error) {
	data, err := report.FromBytes(input)
	if err != nil {
		return nil, err
	}

	r := codequality.ConvertFrom(data)

	output, err := r.Json()
	if err != nil {
		return nil, err
	}

	return output, nil
}

var CodeQuality = codeQualityConverter{}
