package sarifreport

import "github.com/owenrumney/go-sarif/v2/sarif"

type RuleWrapper struct {
	rule *sarif.ReportingDescriptor
}

func (r RuleWrapper) DefaultLevel() string {
	c := r.rule.DefaultConfiguration
	if c == nil {
		return ""
	}

	return r.rule.DefaultConfiguration.Level
}
