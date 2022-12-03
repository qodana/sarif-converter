package converter

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	fixtures := fixture.NewFixtures("../testing/fixtures")

	report, _ := GetConverter("codequality").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, fixtures.ActualJson(), string(report))
}
