package level

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/invocation"
	"sarif-converter/sarifreport/kind"
	"sarif-converter/sarifreport/rule"
)

func GetLevel(result *sarif.Result, invocations invocation.Wrappers, rules rule.Wrappers) string {
	if result.Level != nil {
		return *result.Level
	}

	if kind.GetKind(result) == "fail" {
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
