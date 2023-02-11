package command

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertReSharperSarif(t *testing.T) {
	target := newWrapper()

	target.convert([]string{
		"sarif-converter",
		"--type=codequality",
		"resharper.sarif",
		"output.json",
	})

	assert.Equal(t, fixture.ReSharperCodeQuality(), target.readText("output.json"))
}
