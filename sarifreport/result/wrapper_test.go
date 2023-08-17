package result

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"sarif-converter/sarifreport/invocation"
	"sarif-converter/sarifreport/rule"
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

func TestWrapper_RequireReport_Pass(t *testing.T) {
	target := testingWrapperByKind(p("pass"))

	assert.False(t, target.RequireReport())
}

func TestWrapper_RequireReport_Open(t *testing.T) {
	target := testingWrapperByKind(p("open"))

	assert.True(t, target.RequireReport())
}

func TestWrapper_RequireReport_Informational(t *testing.T) {
	target := testingWrapperByKind(p("informational"))

	assert.False(t, target.RequireReport())
}

func TestWrapper_RequireReport_NotApplicable(t *testing.T) {
	target := testingWrapperByKind(p("notApplicable"))

	assert.False(t, target.RequireReport())
}

func TestWrapper_RequireReport_Review(t *testing.T) {
	target := testingWrapperByKind(p("review"))

	assert.True(t, target.RequireReport())
}

func TestWrapper_RequireReport_Fail(t *testing.T) {
	target := testingWrapperByKind(p("fail"))

	assert.True(t, target.RequireReport())
}

func TestWrapper_RequireReport_Absent(t *testing.T) {
	target := testingWrapperByKind(nil)

	assert.True(t, target.RequireReport())
}

func testingWrapperByKind(kind *string) Wrapper {
	return NewWrapper(
		&sarif.Result{Kind: kind},
		invocation.EmptyWrappers(),
		rule.EmptyWrappers(),
	)
}

func p(s string) *string {
	return &s
}
