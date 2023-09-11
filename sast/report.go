package sast

import (
	"bytes"
	"encoding/json"
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
	filtered := report.OnlyRequireReport()

	sast, err := transformToGLSASTReport(filtered)
	if err != nil {
		return nil, err
	}

	return sast.overrideScan(filtered, time, metadata)
}

func (r Report) overrideScan(report *wrapper.Wrapper, time *now.TimeProvider, metadata meta.Metadata) (*Report, error) {
	overrider := sarif.NewReport(report.Value())
	if time != nil {
		overrider = overrider.WithTimeProvider(time)
	}

	sast := *r.sast
	overrides := overrider.OverrideScan(sast.Scan, metadata)
	overrides.Type = "sast"
	sast.Scan = overrides

	return &Report{sast: &sast}, nil
}

func transformToGLSASTReport(input *wrapper.Wrapper) (*Report, error) {
	b, err := input.Bytes()
	if err != nil {
		return nil, err
	}

	gf := newGitLabFeatures()

	gf.unset()
	sast, err := report.TransformToGLSASTReport(bytes.NewReader(b), "", "", report.Scanner{})
	gf.restore()

	return &Report{sast: sast}, err
}
