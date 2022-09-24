package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/sarif"
)

type SarifRunWrapper struct {
	run *sarif.Run
}

func (r *SarifRunWrapper) CodeQualityElements() []codequality.CodeQualityElement {
	var elements []codequality.CodeQualityElement

	for _, result := range r.results() {
		elements = append(elements, result.CodeQualityElement())
	}

	return elements
}

func (r *SarifRunWrapper) FindRule(id string) RuleWrapper {
	rule, _ := r.run.GetRuleById(id)
	return RuleWrapper{rule: rule}
}

func (r *SarifRunWrapper) results() []SarifResultWrapper {
	var elements []SarifResultWrapper

	for _, result := range r.run.Results {
		elements = append(elements, SarifResultWrapper{result: result, run: r})
	}

	return elements
}
