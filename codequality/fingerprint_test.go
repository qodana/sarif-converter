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

	assert.Equal(t, "32f1254f9dd83b6e3c0bb4609fe0c2b6c5629c29d1061e89dfe0741e70cf8e91", actual)
}

func p(s string) *string {
	return &s
}
