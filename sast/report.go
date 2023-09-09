package sast

import (
	"bytes"
	"encoding/json"
	sarif2 "github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/meta"
	"sarif-converter/now"
	wrapper "sarif-converter/sarifreport/report"
	"sarif-converter/sast/sarif"
)

type Report struct {
	sast *report.Report
}

func (r Report) Json() ([]byte, error) {
	return json.MarshalIndent(r.sast, "", "  ")
}

func ConvertFrom(report *wrapper.Wrapper, time *now.TimeProvider, metadata meta.Metadata) (*Report, error) {
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

	sast.Scan = overrideScan(original, time, sast, metadata)

	return NewReport(sast), nil
}

func overrideScan(original *sarif2.Report, time *now.TimeProvider, sast *report.Report, metadata meta.Metadata) report.Scan {
	r := sarif.NewReport(original)
	if time != nil {
		r = r.WithTimeProvider(time)
	}
	overrides := r.OverrideScan(sast.Scan, metadata)
	overrides.Type = "sast"
	return overrides
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
