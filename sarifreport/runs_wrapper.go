package sarifreport

import (
	"codequality-converter/codequality"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type SarifRunsWrapper struct {
	target []*sarif.Run
}

func (w SarifRunsWrapper) CodeQualityElements() []codequality.Element {
	//goland:noinspection GoPreferNilSlice
	result := []codequality.Element{}

	for _, run := range w.runs() {
		for _, element := range run.CodeQualityElements() {
			result = append(result, element)
		}
	}

	return result
}

func (w SarifRunsWrapper) runs() []SarifRunWrapper {
	//goland:noinspection ALL
	result := []SarifRunWrapper{}

	for _, run := range w.target {
		result = append(result, newSarifRunWrapper(run))
	}

	return result
}

func newSarifRunsWrapper(runs []*sarif.Run) SarifRunsWrapper {
	return SarifRunsWrapper{
		target: runs,
	}
}
