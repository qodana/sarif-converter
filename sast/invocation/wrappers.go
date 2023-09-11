package invocation

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"time"
)

type Wrappers struct {
	invocations []*sarif.Invocation
}

func (w Wrappers) StartTimeUTC() *time.Time {
	times := newTimes(w.invocations, func(invocation *sarif.Invocation) *time.Time {
		return invocation.StartTimeUTC
	})
	return times.Minimal()
}

func (w Wrappers) EndTimeUTC() *time.Time {
	times := newTimes(w.invocations, func(invocation *sarif.Invocation) *time.Time {
		return invocation.EndTimeUTC
	})
	return times.Maximum()
}

func NewWrappers(invocations []*sarif.Invocation) Wrappers {
	return Wrappers{invocations: invocations}
}
