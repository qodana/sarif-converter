package sast

import (
	"bytes"
	"encoding/json"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	sarif "sarif-converter/sarifreport/report"
)

type Report struct {
	sast *report.Report
}

func (r Report) Json() ([]byte, error) {
	return json.MarshalIndent(r.sast, "", "  ")
}

func ConvertFrom(report *sarif.Wrapper) (*Report, error) {
	original := report.Value()
	filtered := report.OnlyRequireReport()
	filteredBytes, err := filtered.Bytes()
	if err != nil {
		return nil, err
	}

	sast, err := transformToGLSASTReport(filteredBytes, err)
	if err != nil {
		return nil, err
	}

	sast.Scan.Scanner.ID = original.Runs[0].Tool.Driver.Name
	sast.Scan.Scanner.Name = original.Runs[0].Tool.Driver.Name
	sast.Scan.Type = "sast"

	return NewReport(sast), nil
}

func NewReport(sast *report.Report) *Report {
	return &Report{sast: sast}
}

func transformToGLSASTReport(input []byte, err error) (*report.Report, error) {
	gf := newGitLabFeatures()

	gf.unset()
	sast, err := report.TransformToGLSASTReport(bytes.NewReader(input), "", "", report.Scanner{})
	gf.restore()

	return sast, err
}
