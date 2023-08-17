package run

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/invocation"
	"sarif-converter/sarifreport/result"
	"sarif-converter/sarifreport/rule"
)

type Wrapper struct {
	Results result.Wrappers
}

func newWrapper(run *sarif.Run) Wrapper {
	invocations := invocation.NewWrappers(run)
	rules := rule.NewWrappers(run)

	return Wrapper{Results: result.NewWrappers(run.Results, invocations, rules)}
}
