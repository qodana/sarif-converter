package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/sarif"
)

type ReportWrapper struct {
	Sarif *sarif.Report
	run   SarifRunWrapper
}

func (r *ReportWrapper) CodeQualityElements() []codequality.CodeQualityElement {
	return r.run.CodeQualityElements()
}

func FromString(content string) (*ReportWrapper, error) {
	data, err := sarif.FromString(content)
	if err != nil {
		return nil, err
	}

	return &ReportWrapper{
		Sarif: data,
		run:   SarifRunWrapper{run: data.Runs[0]},
	}, nil
}
