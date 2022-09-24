package sarifreport

import "github.com/owenrumney/go-sarif/sarif"

type RuleWrapper struct {
	rule *sarif.ReportingDescriptor
}

func (r RuleWrapper) DefaultLevel() interface{} {
	return r.rule.DefaultConfiguration.Level
}
