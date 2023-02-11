package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelResultOnly(t *testing.T) {
	level := "error"

	result := makeWrapper(makeResult("rule1", &level),
		makeRule("rule1", ""))

	assert.Equal(t, level, result.Level())
}

func TestLevelNone(t *testing.T) {
	result := makeWrapper(makeResult("rule1", nil),
		makeRule("rule1", ""))

	assert.Equal(t, "none", result.Level())
}

func TestLevelRuleOnly(t *testing.T) {
	level := "note"

	result := makeWrapper(makeResult("rule1", nil),
		makeRule("rule1", level))

	assert.Equal(t, level, result.Level())
}

func makeWrapper(result sarif.Result, rule sarif.ReportingDescriptor) Issue {
	var wrapper = makeRunWrapper(rule)

	return Issue{
		result: &result,
		run:    wrapper,
	}
}

func makeResult(ruleId string, level *string) sarif.Result {
	return sarif.Result{
		RuleID: &ruleId,
		Level:  level,
	}
}

func makeRunWrapper(rule sarif.ReportingDescriptor) SarifRunWrapper {
	run := sarif.Run{
		Tool: sarif.Tool{
			Driver: &sarif.ToolComponent{
				Rules: []*sarif.ReportingDescriptor{&rule},
			},
		},
	}
	return SarifRunWrapper{run: &run}
}

func makeRule(ruleId string, level string) sarif.ReportingDescriptor {
	return sarif.ReportingDescriptor{
		ID: ruleId,
		DefaultConfiguration: &sarif.ReportingConfiguration{
			Level: level,
		},
	}
}
