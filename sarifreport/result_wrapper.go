package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/sarif"
)

type SarifResultWrapper struct {
	result *sarif.Result
	run    *SarifRunWrapper
}

func (r *SarifResultWrapper) CodeQualityElement() codequality.CodeQualityElement {
	return codequality.CodeQualityElement{
		Description: r.description(),
		Fingerprint: r.fingerprint(),
		Severity:    r.severity(),
		Location:    r.location(),
	}
}

func (r *SarifResultWrapper) severity() string {
	return NewLevel(r).Severity()
}

func (r *SarifResultWrapper) fingerprint() string {
	return "abc"
}

func (r *SarifResultWrapper) location() codequality.CodeQualityLocation {
	return codequality.CodeQualityLocation{
		Path:  r.path(),
		Lines: r.lines(),
	}
}

func (r *SarifResultWrapper) lines() codequality.CodeQualityLocationLine {
	return codequality.CodeQualityLocationLine{
		Begin: *r.result.Locations[0].PhysicalLocation.Region.StartLine,
	}
}

func (r *SarifResultWrapper) path() *string {
	return r.result.Locations[0].PhysicalLocation.ArtifactLocation.URI
}

func (r *SarifResultWrapper) description() *string {
	return r.result.Message.Text
}

func (r *SarifResultWrapper) rule() RuleWrapper {
	return r.run.FindRule(*r.result.RuleID)
}
