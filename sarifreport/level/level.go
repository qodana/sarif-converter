package level

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

func GetLevel(result *sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) string {
	if result.Level != nil {
		return *result.Level
	}

	if kindIsFail(result) {
		r := rules.Find(result)

		configuration := invocations.Find(result).FindConfiguration(r.ID())
		if configuration != nil {
			return configuration.Level
		}

		level := r.DefaultLevel()
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
