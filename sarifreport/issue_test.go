package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelResultOnly(t *testing.T) {
	result := makeWrapper(
		sarif.Result{
			RuleID: ruleId("rule1"),
			Level:  level("error"),
		},
		sarif.ReportingDescriptor{
			ID: "rule1",
		})

	assert.Equal(t, "error", result.Level())
}

func TestLevelNone(t *testing.T) {
	result := makeWrapper(
		sarif.Result{
			RuleID: ruleId("rule1"),
			Level:  nil,
		},
		sarif.ReportingDescriptor{
			ID: "rule1",
		})

	assert.Equal(t, "none", result.Level())
}

func TestLevelRuleOnly(t *testing.T) {
	result := makeWrapper(
		sarif.Result{
			RuleID: ruleId("rule1"),
			Level:  nil,
		},
		sarif.ReportingDescriptor{
			ID:                   "rule1",
			DefaultConfiguration: configuration("note"),
		})

	assert.Equal(t, "note", result.Level())
}

func TestMessage(t *testing.T) {
	result := makeWrapper(
		sarif.Result{
			RuleID:  p("rule1"),
			Message: messageText("message"),
		},
		sarif.ReportingDescriptor{
			ID:              "rule1",
			FullDescription: description("full description"),
		})

	assert.Equal(t, "message", *result.Message())
}

func TestMessageFromRule(t *testing.T) {
	result := makeWrapper(
		sarif.Result{
			RuleID:  ruleId("rule1"),
			Message: messageNoText(),
		},
		sarif.ReportingDescriptor{
			ID:              "rule1",
			FullDescription: description("short description"),
		})

	assert.Equal(t, "short description", *result.Message())
}

func makeWrapper(result sarif.Result, rule sarif.ReportingDescriptor) Issue {
	var wrapper = makeRunWrapper(rule)

	return Issue{
		result: &result,
		run:    wrapper,
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

func ruleId(value string) *string {
	return &value
}

func level(value string) *string {
	return &value
}

func messageText(value string) sarif.Message {
	return sarif.Message{Text: p(value)}
}

func messageNoText() sarif.Message {
	return sarif.Message{}
}

func description(value string) *sarif.MultiformatMessageString {
	return &sarif.MultiformatMessageString{Text: p(value)}
}

func configuration(level string) *sarif.ReportingConfiguration {
	return &sarif.ReportingConfiguration{
		Level: level,
	}
}

func p(s string) *string {
	return &s
}
