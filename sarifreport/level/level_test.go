package level

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLevel(t *testing.T) {
	actual := GetLevel(&sarif.Result{
		Level: p("warning"),
	}, fakeInvocations(), fakeRules())

	assert.Equal(t, "warning", actual)
}

func TestGetLevel_NoLevel(t *testing.T) {
	actual := GetLevel(&sarif.Result{}, fakeInvocations(), fakeRules())

	assert.Equal(t, "warning", actual)
}

func TestGetLevel_NoLabelAndKindIsPass(t *testing.T) {
	actual := GetLevel(&sarif.Result{Kind: p("pass")}, fakeInvocations(), fakeRules())

	assert.Equal(t, "none", actual)
}

func TestGetLevel_NoLabelAndKindIsPassAndHasRule(t *testing.T) {
	actual := GetLevel(&sarif.Result{
		Kind:   p("pass"),
		RuleID: p("error1"),
	}, fakeInvocations(), fakeRules())

	assert.Equal(t, "none", actual)
}

func TestGetLevel_KindIsFailAndNoLevel(t *testing.T) {
	actual := GetLevel(&sarif.Result{
		Kind: p("fail"),
	}, fakeInvocations(), fakeRules())

	assert.Equal(t, "warning", actual)
}

func TestGetLevelKindIsFailAndHasDefaultConfiguration(t *testing.T) {
	actual := GetLevel(&sarif.Result{
		Kind:   p("fail"),
		RuleID: p("invocation2"),
		Provenance: &sarif.ResultProvenance{
			InvocationIndex: pi(1),
		},
	}, fakeInvocations(), fakeRules())

	assert.Equal(t, "note", actual)
}

func fakeRun() *sarif.Run {
	run := sarif.NewRun(*sarif.NewTool(&sarif.ToolComponent{}))

	addInvocation(run, "invocation1", "error")
	addInvocation(run, "invocation2", "note")

	addRuleWithLevel(run, "invocation1", "none")
	addRuleWithLevel(run, "invocation2", "none")

	addRuleWithLevel(run, "error1", "error")
	addRuleWithLevel(run, "warning1", "warning")
	addRuleWithLevel(run, "note1", "note")
	addRuleWithLevel(run, "none1", "none")
	run.AddRule("noLevel1")

	return run
}

func fakeInvocations() invocation.Wrappers {
	return invocation.NewWrappers(fakeRun())
}

func fakeRules() rule.Wrappers {
	return rule.NewWrappers(fakeRun())
}

func addRuleWithLevel(run *sarif.Run, ruleId string, level string) {
	r := run.AddRule(ruleId)
	r.DefaultConfiguration = &sarif.ReportingConfiguration{
		Level: level,
	}
}

func addInvocation(run *sarif.Run, descriptorId string, level string) {
	i := sarif.NewInvocation()
	override := sarif.NewConfigurationOverride()
	override.Descriptor = newReportingDescriptorReference(descriptorId)
	override.Configuration = newReportingConfiguration(level)
	i.AddRuleConfigurationOverride(override)
	run.AddInvocations(i)
}

func newReportingConfiguration(level string) *sarif.ReportingConfiguration {
	configuration := sarif.NewReportingConfiguration()
	configuration.Level = level
	return configuration
}

func newReportingDescriptorReference(id string) *sarif.ReportingDescriptorReference {
	reference := sarif.NewReportingDescriptorReference()
	reference.Id = &id
	return reference
}
