package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToHtml(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := GetConverter("html").Convert(fixtures.SemgrepSarif())
	file, _ := os.Create("sarif-report.html")
	file.Write(report)

	assert.Equal(t, fixture.Html(), string(report))
}
