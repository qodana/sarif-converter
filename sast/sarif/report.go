package sarif

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/meta"
	"sarif-converter/now"
	"sarif-converter/sast/analyzer"
	"sarif-converter/sast/scanner"
	"sarif-converter/sast/scanning"
)

type Report struct {
	Scanner  scanner.Scanner
	Analyzer analyzer.Analyzer
	Scanning scanning.Scanning
}

func (r Report) OverrideScan(scan report.Scan, metadata meta.Metadata) report.Scan {
	scan.Scanner = r.Scanner.ToSast()
	scan.Analyzer = r.Analyzer.ToSast(metadata)
	return r.Scanning.Override(scan)
}

func (r *Report) WithTimeProvider(time now.TimeProvider) *Report {
	r.Scanning = r.Scanning.WithTimeProvider(time)
	return r
}

func NewReport(report *sarif.Report) *Report {
	return &Report{
		Scanning: scanning.NewScanning(report),
		Scanner:  scanner.NewScannerFrom(report),
		Analyzer: analyzer.NewAnalyzer(),
	}
}
