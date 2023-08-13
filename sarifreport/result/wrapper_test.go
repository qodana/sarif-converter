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

	target := NewWrapper(
		&sarif.Result{Kind: p("warning"), Level: p("note")},
		invocations,
		rules,
	)

	assert.Equal(t, "note", target.Level())
}

func p(s string) *string {
	return &s
}