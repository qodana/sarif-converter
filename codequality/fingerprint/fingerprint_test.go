package fingerprint

import (
	"codequality-converter/codequality/element"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFingerprint(t *testing.T) {
	actual := Fingerprint(element.Element{
		Description: p("description"),
		Severity:    "Warning",
		Location: element.Location{
			Path: p("file.txt"),
			Lines: &element.LocationLine{
				Begin: 15,
			},
		},
	})

	assert.Equal(t, "f42130385113fe8a11be6c5d542d99f62ef04c44155294c4e38ca005be422688", actual)
}

func p(s string) *string {
	return &s
}
