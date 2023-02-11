package codequality

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFingerprint(t *testing.T) {
	actual := Fingerprint(Element{
		Description: p("description"),
		Severity:    "Warning",
		Location: Location{
			Path: p("file.txt"),
			Lines: LocationLine{
				Begin: 15,
			},
		},
	})

	assert.Equal(t, "f42130385113fe8a11be6c5d542d99f62ef04c44155294c4e38ca005be422688", actual)
}

func p(s string) *string {
	return &s
}
