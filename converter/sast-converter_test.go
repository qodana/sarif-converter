package converter

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToSast(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	actual, _ := ConvertToSast(fixtures.SemgrepSarif())

	assert.Equal(t, string(fixture.Sast()), string(actual))
}
