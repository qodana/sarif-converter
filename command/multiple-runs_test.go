package command

import (
	"github.com/stretchr/testify/assert"
	"sarif-converter/testing/fixture"
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
