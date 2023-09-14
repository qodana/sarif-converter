package openscap

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sarif-converter/converter"
	"sarif-converter/meta"
	"sarif-converter/now"
	"testing"
)

// https://www.jetbrains.com/qodana/

//go:embed qodana.sarif.json
var sarif []byte

//go:embed gl-sast-report.json
var sast string

//go:embed gl-codequality-report.json
var codequality string

func TestConvertToSastFromQodana(t *testing.T) {
	report, _ := getConverter("sast").Convert(sarif)

	assert.Equal(t, sast, string(report))
}

func TestConvertToCodeQualityFromOpenScap(t *testing.T) {
	report, _ := getConverter("codequality").Convert(sarif)

	assert.Equal(t, codequality, string(report))
}

func getConverter(converterType string) converter.Converter {
	metadata := meta.NewMetadata("0.5.1", "a9323")
	provider := now.NewFakeTime(now.Parse("2023-08-31T15:00:42Z"))

	return converter.GetConverter(converterType, &provider, metadata)
}
