package converter

import (
	"codequality-converter/sarifreport"
	"encoding/json"
	"github.com/owenrumney/go-sarif/sarif"
)

type codeQualityConverter struct {
}

func (c codeQualityConverter) Type() string {
	return "codequality"
}

func (c codeQualityConverter) Convert(input []byte) ([]byte, error) {
	data, err := sarif.FromBytes(input)
	if err != nil {
		return nil, err
	}

	report := sarifreport.NewReport(data)

	output, err := json.MarshalIndent(report.CodeQualityElements(), "", "  ")
	if err != nil {
		return nil, err
	}

	return output, nil
}

var CodeQuality = codeQualityConverter{}
