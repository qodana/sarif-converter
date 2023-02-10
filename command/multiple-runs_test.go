package command

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertMultipleRuns(t *testing.T) {
	target := newWrapper()

	target.convert([]string{
		"sarif-converter",
		"--type=codequality",
		"multi-run.sarif",
		"output.json",
	})

	assert.Equal(t, fixture.MultiRunCodeQuality(), target.readText("output.json"))
}
