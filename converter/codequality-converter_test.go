package converter

import (
	"encoding/json"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"sarif-converter/codequality/element"
	"sarif-converter/filter"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvert(t *testing.T) {
	fixtures := fixture.NewFixtures("../testing/fixtures")

	report, _ := GetConverter("codequality").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, fixtures.ActualJson(), string(report))
}

func TestConvertFromReShaperInspectCodeNoInspections(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.ReSharperNoInspectionsSarif())

	assert.Equal(t, "[]", string(report))
}

func TestConvertFromSecurityCodeScan(t *testing.T) {
	bytes := convertToRelativePath(fixture.SecurityCodeScan(), "file:///home/masakura/tmp/sc")
	report, _ := GetConverter("codequality").Convert(bytes)

	result := codeQuality(report)

	assert.Equal(t, "Controllers/HomeController.cs", *result[0].Location.Path)
}

func TestConvertFromEslintSarif(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.Eslint())

	result := codeQuality(report)

	assert.Equal(t, "eval with argument of type Identifier", *result[0].Description)
}

func TestConvertFromBinSkimSarif(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.BinSkim())

	result := codeQuality(report)

	assert.Nil(t, result[0].Location.Lines)
	assert.Equal(t, "Application code should be compiled with the Spectre mitigations switch (/Qspectre) and toolsets that support it.", *result[0].Description)
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

func codeQuality(report []byte) []element.Element {
	var result []element.Element
	err := json.Unmarshal(report, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func convertToRelativePath(input []byte, srcRoot string) []byte {
	s := toSarif(input)

	rpf := filter.NewRelativePathFilter(srcRoot)
	rpf.Run(s)
	bytes := sarifToBytes(s)
	return bytes
}
