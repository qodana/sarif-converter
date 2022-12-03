package converter

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToSast(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := GetConverter("sast").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, string(fixture.Sast()), string(report))
}
