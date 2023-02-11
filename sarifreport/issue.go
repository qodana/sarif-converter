package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
)

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
	// https://docs.oasis-open.org/sarif/sarif/v2.0/csprd02/sarif-v2.0-csprd02.html#_Toc10127839

	if i.result.Level != nil {
		return *i.result.Level
	}

	var d = i.rule().DefaultLevel()
	if d != "" {
		return d
	}

	return "none"
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
