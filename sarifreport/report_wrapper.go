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

func NewReport(sarif *sarif.Report) ReportWrapper {
	return ReportWrapper{
		Sarif: sarif,
		run:   SarifRunWrapper{run: sarif.Runs[0]},
	}
}
