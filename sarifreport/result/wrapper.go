package result

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/level"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrapper struct {
	result *sarif.Result
}

func (w Wrapper) Level(invocations invocation.Wrappers, runs rule.Wrappers) string {
	return level.GetLevel(w.result, invocations, runs)
}

func newWrapper(result *sarif.Result) Wrapper {
	return Wrapper{result: result}
}
