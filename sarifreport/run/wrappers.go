package run

import (
	"codequality-converter/sarifreport/result"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrappers struct {
	runs []Wrapper
}

func (w Wrappers) Wrappers() result.Wrappers {
	list := result.EmptyWrappers()

	for _, run := range w.runs {
		list = list.Append(run.Results)
	}
	return list
}

func NewWrappers(report *sarif.Report) Wrappers {
	list := make([]Wrapper, 0)

	for _, run := range report.Runs {
		list = append(list, newWrapper(run))
	}

	return Wrappers{runs: list}
}
