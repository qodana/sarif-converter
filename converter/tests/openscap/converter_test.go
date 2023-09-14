package openscap

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sarif-converter/converter/codequality"
	"sarif-converter/converter/sast"
	"testing"
)

// https://github.com/candrews/jumpstart/blob/master/.gitlab-ci.yml
//
// oscap-docker image "${IMAGE_NAME}" xccdf eval --verbose ERROR --fetch-remote-resources --profile "${profile}" --results "openscap-report.xml" --report "openscap-report.html" "${ssgdir}/${ssg}"
// saf convert xccdf_results2hdf -i "openscap-report.xml" -o openscap-report.hdf
// DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1 sarif-multitool convert -t Hdf -o openscap-report.sarif openscap-report.hdf.json
// ./sarif-converter --type sast openscap-report.sarif gl-sast-report.json

//go:embed openscap-report.sarif
var sarif []byte

//go:embed gl-sast-report.json
var sastReport string

//go:embed gl-codequality-report.json
var codequalityReport string

func TestConvertToSastFromOpenScap(t *testing.T) {
	report, _ := sast.NewSastConverterForTest().Convert(sarif)

	assert.Equal(t, sastReport, string(report))
}

func TestConvertToCodeQualityFromOpenScap(t *testing.T) {
	report, _ := codequality.NewCodeQualityConverter().Convert(sarif)

	assert.Equal(t, codequalityReport, string(report))
}
