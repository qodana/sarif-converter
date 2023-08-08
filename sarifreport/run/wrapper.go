package run

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/result"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrapper struct {
	Results result.Wrappers
}

func newWrapper(run *sarif.Run) Wrapper {
	invocations := invocation.NewWrappers(run)
	rules := rule.NewWrappers(run)

	return Wrapper{Results: result.NewWrappers(run.Results, invocations, rules)}
}
