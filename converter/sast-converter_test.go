package converter

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/meta"
	"sarif-converter/now"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToSast(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := newSastConverter().Convert(fixtures.SemgrepSarif())

	assert.Equal(t, string(fixture.Sast()), string(report))
}

func TestConvert_SastMetadata(t *testing.T) {
	fixtures := fixture.NewFixtures("").Sast.Metadata

	report, _ := newSastConverter().Convert(fixtures.Sarif())

	assert.Equal(t, fixtures.Sast(), string(report))
}

func newSastConverter() Converter {
	provider := now.NewFakeTime(now.Parse("2023-08-31T15:00:42Z"))
	return GetConverter("sast", &provider, meta.NewMetadata("0.5.1", "a9323"))
}
