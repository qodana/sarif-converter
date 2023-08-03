package rule

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	target := newTarget()

	rule := target.Find(&sarif.Result{RuleID: p("rule2")})

	assert.Equal(t, "RULE2", *rule.name())
}

func TestFind_NotFound(t *testing.T) {
	target := newTarget()

	rule := target.Find(&sarif.Result{RuleID: p("not-exists")})

	assert.False(t, rule.exists())
}

func TestFind_NoRuleInResult(t *testing.T) {
	target := newTarget()

	rule := target.Find(&sarif.Result{})

	assert.False(t, rule.exists())
}

func newTarget() Wrappers {
	target := NewWrappers(
		&sarif.Run{
			Tool: sarif.Tool{
				Driver: &sarif.ToolComponent{
					Rules: []*sarif.ReportingDescriptor{
						{ID: "rule1", Name: p("RULE1")},
						{ID: "rule2", Name: p("RULE2")},
					},
				},
			},
		})
	return target
}

func p(s string) *string {
	return &s
}
