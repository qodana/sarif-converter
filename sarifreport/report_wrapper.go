package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type ReportWrapper struct {
	runs SarifRunsWrapper
}

func (r *ReportWrapper) Issues() []Issue {
	return r.runs.issues()
}

func NewReport(sarif *sarif.Report) *ReportWrapper {
	return &ReportWrapper{
		runs: newSarifRunsWrapper(sarif.Runs),
	}
}

func FromBytes(data []byte) (*ReportWrapper, error) {
	report, err := sarif.FromBytes(data)
	if err != nil {
		return nil, err
	}
	return NewReport(report), nil
}
