package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/sarif"
)

type ResultWrapper struct {
	result *sarif.Result
	run    *SarifRunWrapper
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
	return NewLevel(r).Severity()
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

func (r *ResultWrapper) path() *string {
	return r.result.Locations[0].PhysicalLocation.ArtifactLocation.URI
}

func (r *ResultWrapper) description() *string {
	return r.result.Message.Text
}

func (r *ResultWrapper) rule() RuleWrapper {
	return r.run.FindRule(*r.result.RuleID)
}
