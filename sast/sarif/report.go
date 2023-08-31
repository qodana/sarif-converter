package sarif

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sast/analyzer"
	"sarif-converter/sast/scanner"
)

type Report struct {
	Scanner  scanner.Scanner
	Analyzer analyzer.Analyzer
}

func FromBytes(input []byte) (*Report, error) {
	report, err := sarif.FromBytes(input)
	if err != nil {
		return nil, err
	}

	return &Report{
		Scanner:  scanner.NewScanner(report.Runs[0].Tool.Driver),
		Analyzer: analyzer.NewAnalyzer(),
	}, nil
}
