package converter

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToSast(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := GetConverter("sast").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, string(fixture.Sast()), string(report))
}
