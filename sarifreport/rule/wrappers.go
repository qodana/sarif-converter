package rule

import "github.com/owenrumney/go-sarif/v2/sarif"

type Wrappers struct {
	rules []Wrapper
	run   *sarif.Run
}

func (w Wrappers) Find(result *sarif.Result) Wrapper {
	if result.RuleID == nil {
		return newWrapper(nil)
	}

	rule, _ := w.run.GetRuleById(*result.RuleID)
	return newWrapper(rule)
}

func NewWrappers(run *sarif.Run) Wrappers {
	return Wrappers{run: run}
}
