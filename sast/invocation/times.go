package invocation

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"time"
)

type times struct {
	times []time.Time
}

func (t times) Minimal() *time.Time {
	return t.extract(func(left time.Time, right time.Time) bool {
		return left.Before(right)
	})
}

func (t times) Maximum() *time.Time {
	return t.extract(func(left time.Time, right time.Time) bool {
		return left.After(right)
	})
}

func (t times) extract(compare func(left time.Time, right time.Time) bool) *time.Time {
	var value *time.Time

	for _, i := range t.times {
		if value == nil {
			value = &i
			continue
		}
		if compare(*value, i) {
			value = &i
		}
	}
	return value
}

func newTimes(invocations []*sarif.Invocation, getTime func(invocation *sarif.Invocation) *time.Time) times {
	list := make([]time.Time, 0)

	for _, invocation := range invocations {
		t := getTime(invocation)
		if t != nil {
			list = append(list, *t)
		}
	}

	return times{times: list}
}
