package converter

import (
	"codequality-converter/sarifreport"
	"encoding/json"
	"github.com/owenrumney/go-sarif/sarif"
)

func Convert(input []byte) ([]byte, error) {
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
