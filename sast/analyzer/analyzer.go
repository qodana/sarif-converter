package analyzer

import (
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/meta"
)

type Analyzer struct {
}

func NewAnalyzer() Analyzer {
	return Analyzer{}
}

func (a Analyzer) ToSast(metadata meta.Metadata) report.AnalyzerDetails {
	return report.AnalyzerDetails{
		ID:      "sarif-converter",
		Name:    "SARIF Converter",
		Version: metadata.Version,
		Vendor:  report.Vendor{Name: "SARIF Converter"},
	}
}
