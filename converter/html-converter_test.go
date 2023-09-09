package converter

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToHtml(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := GetConverter("html").Convert(fixtures.SemgrepSarif())

	assert.Equal(t, fixture.Html(), string(report))
}
