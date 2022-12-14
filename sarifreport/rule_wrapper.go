package sarifreport

import "github.com/owenrumney/go-sarif/sarif"

type RuleWrapper struct {
	rule *sarif.ReportingDescriptor
}

func (r RuleWrapper) DefaultLevel() *string {
	c := r.rule.DefaultConfiguration
	if c == nil {
		return nil
	}

	v := r.rule.DefaultConfiguration.Level
	if v == nil {
		return nil
	}

	l := v.(string)
	return &l
}
