package message

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"sarif-converter/sarifreport/rule"
	"testing"
)

func TestGetTextMessage_Text(t *testing.T) {
	actual := GetTextMessage(&sarif.Result{
		RuleID: p("rule1"),
		Message: sarif.Message{
			Text: p("result.message"),
		},
	}, fakeRules())

	assert.Equal(t, "result.message", actual)
}

func TestGetTextMessage_ByRule_Simple(t *testing.T) {
	actual := GetTextMessage(&sarif.Result{
		RuleID: p("rule1"),
		Message: sarif.Message{
			ID: p("Simple"),
		},
	}, fakeRules())

	assert.Equal(t, "rule.messageStrings.Simple.text", actual)
}

func TestGetTextMessage_ByRule_WithArguments(t *testing.T) {
	actual := GetTextMessage(&sarif.Result{
		RuleID: p("rule1"),
		Message: sarif.Message{
			ID:        p("WithArgs"),
			Arguments: []string{"a", "b"},
		},
	}, fakeRules())

	assert.Equal(t, "a is b", actual)
}

func p(s string) *string {
	return &s
}

func fakeRun() *sarif.Run {
	run := sarif.NewRun(*sarif.NewTool(&sarif.ToolComponent{}))

	addRule(run, "rule1", &sarif.MessageStrings{
		"Simple": sarif.MultiformatMessageString{
			Text: p("rule.messageStrings.Simple.text"),
		},
		"WithArgs": sarif.MultiformatMessageString{
			Text: p("{0} is {1}"),
		},
	})

	return run
}

func addRule(run *sarif.Run, ruleId string, messages *sarif.MessageStrings) {
	rule1 := run.AddRule(ruleId)
	rule1.MessageStrings = messages
}

func fakeRules() rule.Wrappers {
	return rule.NewWrappers(fakeRun())
}
