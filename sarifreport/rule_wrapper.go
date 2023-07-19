package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
)

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

func (r RuleWrapper) Message() *string {
	// TODO Tentative since messageStrings is not implemented in go-sarif
	if r.rule.FullDescription != nil {
		return r.rule.FullDescription.Text
	}
	return nil
}
