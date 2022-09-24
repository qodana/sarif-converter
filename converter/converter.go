package converter

import (
	"codequality-converter/sarifreport"
	"encoding/json"
)

func convert(input string) string {
	sarif, _ := sarifreport.FromString(input)

	output, _ := json.MarshalIndent(sarif.CodeQualityElements(), "", "  ")
	return string(output)
}
