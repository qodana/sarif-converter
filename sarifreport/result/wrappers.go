package result

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrappers struct {
	results []Wrapper
}

func (w Wrappers) Get(i int) Wrapper {
	return w.results[i]
}

func (w Wrappers) Iter() <-chan Wrapper {
	out := make(chan Wrapper, len(w.results))
	go func() {
		defer close(out)
		for _, result := range w.results {
			out <- result
		}
	}()
	return out
}

func (w Wrappers) Append(results Wrappers) Wrappers {
	list := make([]Wrapper, 0)

	for _, result := range w.results {
		list = append(list, result)
	}
	for _, result := range results.results {
		list = append(list, result)
	}
	return Wrappers{results: list}
}

func EmptyWrappers() Wrappers {
	return Wrappers{}
}

func NewWrappers(results []*sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) Wrappers {
	list := make([]Wrapper, 0)

	for _, result := range results {
		list = append(list, newWrapper(result, invocations, rules))
	}
	return Wrappers{results: list}
}
