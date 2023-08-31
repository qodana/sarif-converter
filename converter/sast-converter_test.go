package converter

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/now"
	"sarif-converter/testing/fixture"
	"testing"
	"time"
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
	provider := now.NewFakeTime(parse("2023-08-31T15:00:42Z"))
	return GetConverter("sast", &provider)
}

func parse(value string) time.Time {
	s, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return s
}
