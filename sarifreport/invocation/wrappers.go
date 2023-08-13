package invocation

import "github.com/owenrumney/go-sarif/v2/sarif"

type Wrappers struct {
	invocations []Wrapper
}

func (w Wrappers) Find(result *sarif.Result) Wrapper {
	i := invocationIndex(result)
	if i == nil {
		return newWrapper(nil)
	}

	index := *i

	if index >= len(w.invocations) {
		return newWrapper(nil)
	}
	if index < 0 {
		return newWrapper(nil)
	}

	return w.invocations[index]
}

func NewWrappers(run *sarif.Run) Wrappers {
	invocations := make([]Wrapper, 0)

	for _, element := range run.Invocations {
		invocations = append(invocations, newWrapper(element))
	}

	return Wrappers{
		invocations: invocations,
	}
}

func EmptyWrappers() Wrappers {
	return Wrappers{invocations: make([]Wrapper, 0)}
}

func invocationIndex(result *sarif.Result) *int {
	if result.Provenance == nil {
		return nil
	}
	if result.Provenance.InvocationIndex == nil {
		return nil
	}
	return result.Provenance.InvocationIndex
}
