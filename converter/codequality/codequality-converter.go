package codequality

import (
	"sarif-converter/codequality"
	"sarif-converter/sarifreport/report"
)

var CodeQuality = "codequality"

type CodeQualityConverter struct {
}

func (c CodeQualityConverter) Convert(input []byte) ([]byte, error) {
	data, err := report.FromBytes(input)
	if err != nil {
		return nil, err
	}

	r := codequality.ConvertFrom(data)

	return r.Json()
}

func NewCodeQualityConverter() CodeQualityConverter {
	return CodeQualityConverter{}
}
