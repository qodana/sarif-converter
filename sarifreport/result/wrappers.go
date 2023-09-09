package result

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/invocation"
	"sarif-converter/sarifreport/rule"
)

type Wrappers struct {
	results []Wrapper
	value   []*sarif.Result
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

func (w Wrappers) OnlyRequireReport() Wrappers {
	listWrapper := make([]Wrapper, 0)
	listResult := make([]*sarif.Result, 0)

	for _, result := range w.results {
		if result.RequireReport() {
			listWrapper = append(listWrapper, result)
			listResult = append(listResult, result.Value())
		}
	}

	return Wrappers{results: listWrapper, value: listResult}
}

func (w Wrappers) Value() []*sarif.Result {
	return w.value
}

func EmptyWrappers() Wrappers {
	return Wrappers{}
}

func NewWrappers(results []*sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) Wrappers {
	list := make([]Wrapper, 0)

	for _, result := range results {
		list = append(list, NewWrapper(result, invocations, rules))
	}
	return Wrappers{results: list, value: results}
}
