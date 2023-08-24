package fingerprint

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/codequality/element"
	"testing"
)

func TestFingerprint(t *testing.T) {
	actual := Fingerprint(element.Element{
		CheckName:   p("security/detect-eval-with-expression"),
		Description: p("description"),
		Severity:    "Warning",
		Location: element.Location{
			Path: p("file.txt"),
			Lines: &element.LocationLine{
				Begin: 15,
			},
		},
	})

	assert.Equal(t, "b4ca3486a7f81d2847ccb81fa126a46faac17dc136d4f0dcee26301a9bf9ac1d", actual)
}

func p(s string) *string {
	return &s
}
