package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type SarifRunWrapper struct {
	run *sarif.Run
}

func (r SarifRunWrapper) FindRule(id string) RuleWrapper {
	rule, _ := r.run.GetRuleById(id)
	return RuleWrapper{rule: rule}
}

func (r SarifRunWrapper) issues() []Issue {
	issues := make([]Issue, len(r.run.Results))

	for i, result := range r.run.Results {
		issues[i] = NewIssue(result, r)
	}

	return issues
}

func newSarifRunWrapper(r *sarif.Run) SarifRunWrapper {
	return SarifRunWrapper{
		run: r,
	}
}
