package result

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/invocation"
	"sarif-converter/sarifreport/kind"
	"sarif-converter/sarifreport/level"
	"sarif-converter/sarifreport/rule"
)

type Wrapper struct {
	result      *sarif.Result
	invocations invocation.Wrappers
	rules       rule.Wrappers
}

func (w Wrapper) Level() string {
	return level.GetLevel(w.result, w.invocations, w.rules)
}

func (w Wrapper) TextMessage() *string {
	return w.result.Message.Text
}

func (w Wrapper) FirstLocation() *sarif.Location {
	return w.result.Locations[0]
}

func (w Wrapper) Rule() rule.Wrapper {
	return w.rules.Find(w.result)
}

func (w Wrapper) kind() string {
	return kind.GetKind(w.result)
}

func NewWrapper(result *sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) Wrapper {
	return Wrapper{
		result:      result,
		invocations: invocations,
		rules:       rules,
	}
}
