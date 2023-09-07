package report

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/result"
	"sarif-converter/sarifreport/run"
)

type Wrapper struct {
	runs  run.Wrappers
	value *sarif.Report
}

func (w Wrapper) Results() result.Wrappers {
	return w.runs.Results()
}

func (w Wrapper) OnlyRequireReport() Wrapper {
	runs := w.runs.OnlyRequireReport()

	report := *w.value
	report.Runs = runs.Value()

	return Wrapper{
		runs:  runs,
		value: &report,
	}
}

func (w Wrapper) Value() *sarif.Report {
	return w.value
}

func NewReport(report *sarif.Report) *Wrapper {
	return &Wrapper{
		runs:  run.NewWrappers(report),
		value: report,
	}
}

func FromBytes(data []byte) (*Wrapper, error) {
	report, err := sarif.FromBytes(data)
	if err != nil {
		return nil, err
	}
	return NewReport(report), nil
}
