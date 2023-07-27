package level

import (
	"codequality-converter/sarifreport/invocation"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

func GetLevel(result *sarif.Result, run *sarif.Run) string {
	if result.Level != nil {
		return *result.Level
	}

	if kindIsFail(result) {
		rule := findRule(run, result.RuleID)

		configuration := invocation.NewWrappers(run).Find(result).FindConfiguration(rule)
		if configuration != nil {
			return configuration.Level
		}

		level := defaultLevel(rule)
		if level != nil {
			return *level
		}

		return "warning"
	}

	return "none"
}

func kindIsFail(r *sarif.Result) bool {
	if r.Kind == nil {
		return false
	}
	return *r.Kind == "fail"
}

func findRule(run *sarif.Run, ruleId *string) *sarif.ReportingDescriptor {
	if ruleId == nil {
		return nil
	}

	rule, _ := run.GetRuleById(*ruleId)
	return rule
}

func defaultLevel(rule *sarif.ReportingDescriptor) *string {
	if rule == nil {
		return nil
	}
	if rule.DefaultConfiguration == nil {
		return nil
	}
	return &rule.DefaultConfiguration.Level
}
