package sast

import (
	bytes2 "bytes"
	"encoding/json"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	report2 "sarif-converter/sarifreport/report"
)

type Report struct {
}

func ConvertFrom(report *report2.Wrapper) ([]byte, error) {
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

	bytes, err := json.MarshalIndent(sast, "", "  ")
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func transformToGLSASTReport(input []byte, err error) (*report.Report, error) {
	gf := newGitLabFeatures()

	gf.unset()
	sast, err := report.TransformToGLSASTReport(bytes2.NewReader(input), "", "", report.Scanner{})
	gf.restore()

	return sast, err
}
