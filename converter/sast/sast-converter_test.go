package sast

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToSast(t *testing.T) {
	fixtures := fixture.NewFixtures("")

	report, _ := NewSastConverterForTest().Convert(fixtures.SemgrepSarif())

	assert.Equal(t, string(fixture.Sast()), string(report))
}

func TestConvert_SastMetadata(t *testing.T) {
	fixtures := fixture.NewFixtures("").Sast.Metadata

	report, _ := NewSastConverterForTest().Convert(fixtures.Sarif())

	assert.Equal(t, fixtures.Sast(), string(report))
}
