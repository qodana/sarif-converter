package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/sarif"
)

type ResultWrapper struct {
	result *sarif.Result
	run    *SarifRunWrapper
}

func NewResultWrapper(result *sarif.Result, run *SarifRunWrapper) ResultWrapper {
	return ResultWrapper{
		result: result,
		run:    run,
	}
}

func (r *ResultWrapper) CodeQualityElement() codequality.CodeQualityElement {
	return codequality.CodeQualityElement{
		Description: r.description(),
		Fingerprint: r.fingerprint(),
		Severity:    r.severity(),
		Location:    r.location(),
	}
}

func (r *ResultWrapper) severity() string {
	return NewLevel(*r.level()).Severity()
}

func (r *ResultWrapper) fingerprint() string {
	return Fingerprint(r.location(), r.description())
}

func (r *ResultWrapper) location() codequality.CodeQualityLocation {
	return codequality.CodeQualityLocation{
		Path:  r.path(),
		Lines: r.lines(),
	}
}

func (r *ResultWrapper) lines() codequality.CodeQualityLocationLine {
	return codequality.CodeQualityLocationLine{
		Begin: *r.result.Locations[0].PhysicalLocation.Region.StartLine,
	}
}

func (r *ResultWrapper) level() *string {
	// https://docs.oasis-open.org/sarif/sarif/v2.0/csprd02/sarif-v2.0-csprd02.html#_Toc10127839

	if r.result.Level != nil {
		return r.result.Level
	}

	var d = r.rule().DefaultLevel()
	if d != nil {
		return d
	}

	var defaultLevel = "none"
	return &defaultLevel
}

func (r *ResultWrapper) path() *string {
	return r.result.Locations[0].PhysicalLocation.ArtifactLocation.URI
}

func (r *ResultWrapper) description() *string {
	return r.result.Message.Text
}

func (r *ResultWrapper) rule() RuleWrapper {
	return r.run.FindRule(*r.result.RuleID)
}
