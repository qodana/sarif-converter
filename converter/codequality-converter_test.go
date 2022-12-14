package converter

import (
	"codequality-converter/codequality"
	"codequality-converter/testing/fixture"
	"encoding/json"
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

	var result []codequality.CodeQualityElement
	err := json.Unmarshal(report, &result)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 18, len(result))
}

func TestConvertFromReShaperInspectCodeNoInspections(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.ReSharperNoInspectionsSarif())

	assert.Equal(t, "[]", string(report))
}
