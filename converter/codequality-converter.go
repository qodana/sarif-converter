package converter

import (
	"codequality-converter/codequality"
	"codequality-converter/sarifreport"
)

type codeQualityConverter struct {
}

func (c codeQualityConverter) Type() string {
	return "codequality"
}

func (c codeQualityConverter) Convert(input []byte) ([]byte, error) {
	data, err := sarifreport.FromBytes(input)
	if err != nil {
		return nil, err
	}

	report := codequality.ConvertFrom(*data)

	output, err := report.Json()
	if err != nil {
		return nil, err
	}

	return output, nil
}

var CodeQuality = codeQualityConverter{}
