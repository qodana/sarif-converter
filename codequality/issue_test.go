package codequality

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/result"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation(t *testing.T) {
	target := testingIssue("foo.js", newRegion(10), emptyMessage(), nil)

	assert.Equal(t, Location{
		Path:  p("foo.js"),
		Lines: &LocationLine{Begin: 10},
	}, target.location())
}

func TestLocation_NoRegion(t *testing.T) {
	target := testingIssue("foo.js", nil, emptyMessage(), nil)

	assert.Equal(t, Location{
		Path: p("foo.js"),
	}, target.location())
}

func TestDescription(t *testing.T) {
	target := testingIssue("foo.js", nil, textMessage("original message"), nil)

	assert.Equal(t, "original message", *target.description())
}

func TestDescription_FromRuleFullDescription(t *testing.T) {
	target := testingIssue("foo.js", nil, emptyMessage(), ruleID("rule1"))

	assert.Equal(t, "rule1 full description", *target.description())
}

func testingIssue(file string, region *sarif.Region, message sarif.Message, ruleId *string) issue {
	return newIssue(result.NewWrapper(&sarif.Result{
		Message: message,
		RuleID:  ruleId,
		Locations: []*sarif.Location{
			{PhysicalLocation: &sarif.PhysicalLocation{
				ArtifactLocation: &sarif.ArtifactLocation{
					URI: &file,
				},
				Region: region,
			}},
		},
	},
		invocation.EmptyWrappers(),
		rules(),
	))
}

func emptyMessage() sarif.Message {
	return sarif.Message{}
}

func textMessage(text string) sarif.Message {
	return sarif.Message{
		Text: &text,
	}
}

func ruleID(id string) *string {
	return &id
}

func newRegion(startLine int) *sarif.Region {
	return &sarif.Region{
		StartLine: &startLine,
	}
}

func rules() rule.Wrappers {
	return rule.NewWrappers(&sarif.Run{
		Tool: sarif.Tool{
			Driver: &sarif.ToolComponent{
				Rules: []*sarif.ReportingDescriptor{
					{
						ID: "rule1",
						FullDescription: &sarif.MultiformatMessageString{
							Text: p("rule1 full description"),
						},
					},
				},
			},
		},
	})
}
