package invocation

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	target := newTarget()

	result := sarif.Result{Provenance: &sarif.ResultProvenance{InvocationIndex: pi(1)}}
	invocation := target.Find(&result)

	assert.True(t, invocation.exists())
}

func TestFind_NoInvocationIndex(t *testing.T) {
	target := newTarget()

	result := sarif.Result{Provenance: &sarif.ResultProvenance{}}
	invocation := target.Find(&result)

	assert.False(t, invocation.exists())
}

func TestFind_NoProvenance(t *testing.T) {
	target := newTarget()

	result := sarif.Result{}
	invocation := target.Find(&result)

	assert.False(t, invocation.exists())
}

func TestFind_OutOfRange(t *testing.T) {
	target := newTarget()

	result := sarif.Result{Provenance: &sarif.ResultProvenance{InvocationIndex: pi(2)}}
	invocation := target.Find(&result)

	assert.False(t, invocation.exists())
}

func TestFind_MinusIndex(t *testing.T) {
	target := newTarget()

	result := sarif.Result{Provenance: &sarif.ResultProvenance{InvocationIndex: pi(-1)}}
	invocation := target.Find(&result)

	assert.False(t, invocation.exists())
}

func newTarget() Wrappers {
	return NewWrappers(&sarif.Run{
		Invocations: []*sarif.Invocation{
			{},
			{},
		},
	})
}
