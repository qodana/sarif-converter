package sarifreport

import (
	"codequality-converter/sarifreport/invocation"
	"codequality-converter/sarifreport/result"
	"codequality-converter/sarifreport/rule"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

// TOOD Deprecated
type Issue struct {
	result *sarif.Result
	run    SarifRunWrapper
}

type Location struct {
	Path      *string
	StartLine int
}

func newIssue(result *sarif.Result, run SarifRunWrapper) Issue {
	return Issue{
		result: result,
		run:    run,
	}
}

func (i *Issue) Level() string {
	r := result.NewWrapper(i.result, invocation.NewWrappers(i.run.run), rule.NewWrappers(i.run.run))
	return r.Level()
}

func (i *Issue) path() *string {
	return i.result.Locations[0].PhysicalLocation.ArtifactLocation.URI
}

func (i *Issue) description() *string {
	return i.result.Message.Text
}

func (i *Issue) rule() RuleWrapper {
	return i.run.FindRule(*i.result.RuleID)
}

func (i *Issue) Message() *string {
	return i.result.Message.Text
}

func (i *Issue) Location() Location {
	return Location{
		Path:      i.path(),
		StartLine: *i.result.Locations[0].PhysicalLocation.Region.StartLine,
	}
}
