package report

import (
	"codequality-converter/sarifreport/result"
	"codequality-converter/sarifreport/run"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrapper struct {
	runs run.Wrappers
}

func (w Wrapper) Results() result.Wrappers {
	return w.runs.Results()
}

func NewReport(report *sarif.Report) *Wrapper {
	return &Wrapper{
		runs: run.NewWrappers(report),
	}
}

func FromBytes(data []byte) (*Wrapper, error) {
	report, err := sarif.FromBytes(data)
	if err != nil {
		return nil, err
	}
	return NewReport(report), nil
}
