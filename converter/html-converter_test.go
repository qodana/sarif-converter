package converter

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToHtml(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := ConvertToHtml(fixtures.SemgrepSarif())

	assert.Equal(t, fixture.Html(), string(report))
}
