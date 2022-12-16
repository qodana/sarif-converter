package converter

import (
	"codequality-converter/codequality"
	"codequality-converter/filter"
	"codequality-converter/testing/fixture"
	"encoding/json"
	"github.com/owenrumney/go-sarif/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	fixtures := fixture.NewFixtures("../testing/fixtures")

	report, _ := GetConverter("codequality").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, fixtures.ActualJson(), string(report))
}

func TestConvertFromReSharperInspectCode(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.ReSharperSarif())

	result := codeQuality(report)

	assert.Equal(t, 18, len(result))
}

func TestConvertFromReShaperInspectCodeNoInspections(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.ReSharperNoInspectionsSarif())

	assert.Equal(t, "[]", string(report))
}

func TestConvertFromSecurityCodeScan(t *testing.T) {
	s := toSarif(fixture.SecurityCodeScan())

	rpf := filter.NewRelativePathFilter("file:///home/masakura/tmp/sc")
	rpf.Run(s)
	bytes := sarifToBytes(s)

	report, _ := GetConverter("codequality").Convert(bytes)

	result := codeQuality(report)

	assert.Equal(t, "Controllers/HomeController.cs", *result[0].Location.Path)
}

func sarifToBytes(report *sarif.Report) []byte {
	j, err := json.Marshal(report)
	if err != nil {
		panic(err)
	}
	return j
}

func toSarif(report []byte) *sarif.Report {
	s, err := sarif.FromBytes(report)
	if err != nil {
		panic(err)
	}
	return s
}

func codeQuality(report []byte) []codequality.CodeQualityElement {
	var result []codequality.CodeQualityElement
	err := json.Unmarshal(report, &result)
	if err != nil {
		panic(err)
	}
	return result
}
