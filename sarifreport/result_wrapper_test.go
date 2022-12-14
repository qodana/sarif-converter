package sarifreport

import (
	"github.com/owenrumney/go-sarif/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelResultOnly(t *testing.T) {
	level := "error"

	result := makeWrapper(makeResult("rule1", &level),
		makeRule("rule1", nil))

	assert.Equal(t, level, *result.level())
}

func TestLevelNone(t *testing.T) {
	result := makeWrapper(makeResult("rule1", nil),
		makeRule("rule1", nil))

	assert.Equal(t, "none", *result.level())
}

func TestLevelRuleOnly(t *testing.T) {
	level := "note"

	result := makeWrapper(makeResult("rule1", nil),
		makeRule("rule1", &level))

	assert.Equal(t, level, *result.level())
}

func makeWrapper(result sarif.Result, rule sarif.ReportingDescriptor) ResultWrapper {
	var wrapper = makeRunWrapper(rule)

	return ResultWrapper{
		result: &result,
		run:    &wrapper,
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

func makeRule(ruleId string, level *string) sarif.ReportingDescriptor {
	var v interface{}
	if level == nil {
		v = nil
	} else {
		v = *level
	}

	return sarif.ReportingDescriptor{
		ID: ruleId,
		DefaultConfiguration: &sarif.ReportingConfiguration{
			Level: v,
		},
	}
}
