package command

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertMultipleFiles(t *testing.T) {
	target := newWrapper()

	target.convert([]string{
		"sarif-converter",
		"--type=codequality",
		"semgrep.sarif",
		"security-scan.sarif",
		"output.json",
	})

	assert.Equal(t, fixture.MultiRunCodeQuality(), target.readText("output.json"))
}
