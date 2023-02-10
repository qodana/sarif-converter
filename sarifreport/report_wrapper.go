package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type ReportWrapper struct {
	Sarif *sarif.Report
	runs  SarifRunsWrapper
}

func (r *ReportWrapper) CodeQualityElements() []codequality.CodeQualityElement {
	return r.runs.CodeQualityElements()
}

func NewReport(sarif *sarif.Report) ReportWrapper {
	return ReportWrapper{
		Sarif: sarif,
		runs:  newSarifRunsWrapper(sarif.Runs),
	}
}
