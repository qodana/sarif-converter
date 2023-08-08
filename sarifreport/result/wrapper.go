package result

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/level"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrapper struct {
	result      *sarif.Result
	invocations invocation.Wrappers
	rules       rule.Wrappers
}

func (w Wrapper) Level() string {
	return level.GetLevel(w.result, w.invocations, w.rules)
}

func (w Wrapper) Message() *string {
	return w.result.Message.Text
}

func newWrapper(result *sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) Wrapper {
	return Wrapper{
		result:      result,
		invocations: invocations,
		rules:       rules,
	}
}
