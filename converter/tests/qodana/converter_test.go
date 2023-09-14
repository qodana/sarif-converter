package openscap

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sarif-converter/converter/codequality"
	"sarif-converter/converter/sast"
	"testing"
)

// https://www.jetbrains.com/qodana/

//go:embed qodana.sarif.json
var sarif []byte

//go:embed gl-sast-report.json
var sastReport string

//go:embed gl-codequality-report.json
var codequalityReport string

func TestConvertToSastFromQodana(t *testing.T) {
	report, _ := sast.NewSastConverterForTest().Convert(sarif)

	assert.Equal(t, sastReport, string(report))
}

func TestConvertToCodeQualityFromOpenScap(t *testing.T) {
	report, _ := codequality.NewCodeQualityConverter().Convert(sarif)

	assert.Equal(t, codequalityReport, string(report))
}
