package run

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/result"
)

type Wrappers struct {
	runs  []Wrapper
	value []*sarif.Run
}

func (w Wrappers) Results() result.Wrappers {
	list := result.EmptyWrappers()

	for _, run := range w.runs {
		list = list.Append(run.Results)
	}
	return list
}

func (w Wrappers) OnlyRequireReport() Wrappers {
	length := len(w.runs)
	listRuns := make([]Wrapper, length)
	listValue := make([]*sarif.Run, length)

	for i, run := range w.runs {
		filtered := run.OnlyRequireReport()
		listRuns[i] = filtered
		listValue[i] = filtered.Value()
	}
	return Wrappers{runs: listRuns, value: listValue}
}

func (w Wrappers) Value() []*sarif.Run {
	return w.value
}

func NewWrappers(report *sarif.Report) Wrappers {
	list := make([]Wrapper, 0)

	for _, run := range report.Runs {
		list = append(list, newWrapper(run))
	}

	return Wrappers{runs: list, value: report.Runs}
}
