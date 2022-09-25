package converter

import (
	"codequality-converter/sarifreport"
	"encoding/json"
	"github.com/owenrumney/go-sarif/sarif"
)

func Convert(input []byte) []byte {
	data, _ := sarif.FromBytes(input)
	report := sarifreport.NewReport(data)

	output, _ := json.MarshalIndent(report.CodeQualityElements(), "", "  ")

	return output
}
