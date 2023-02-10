package command

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
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
