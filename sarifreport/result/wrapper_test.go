package result

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel(t *testing.T) {
	r := &sarif.Run{
		Invocations: []*sarif.Invocation{},
		Tool: sarif.Tool{
			Driver: &sarif.ToolComponent{
				Rules: []*sarif.ReportingDescriptor{},
			},
		},
		Results: []*sarif.Result{
			{Kind: p("warning"), Level: p("note")},
		},
	}
	invocations := invocation.NewWrappers(r)
	rules := rule.NewWrappers(r)

	target := newWrapper(&sarif.Result{Kind: p("warning"), Level: p("note")})

	assert.Equal(t, "note", target.Level(invocations, rules))
}

func p(s string) *string {
	return &s
}
