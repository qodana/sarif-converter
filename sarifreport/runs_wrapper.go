package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type SarifRunsWrapper struct {
	target []*sarif.Run
}

func (w SarifRunsWrapper) issues() []Issue {
	result := make([]Issue, 0)

	for _, run := range w.runs() {
		result = append(result, run.issues()...)
	}

	return result
}

func (w SarifRunsWrapper) runs() []SarifRunWrapper {
	result := make([]SarifRunWrapper, len(w.target))

	for i, run := range w.target {
		result[i] = newSarifRunWrapper(run)
	}

	return result
}

func newSarifRunsWrapper(runs []*sarif.Run) SarifRunsWrapper {
	return SarifRunsWrapper{
		target: runs,
	}
}
