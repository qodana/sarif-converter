package openscap

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sarif-converter/converter"
	"sarif-converter/meta"
	"sarif-converter/now"
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
var sast string

//go:embed gl-codequality-report.json
var codequality string

func TestConvertToSastFromOpenScap(t *testing.T) {

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
