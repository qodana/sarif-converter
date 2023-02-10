package command

import (
	"codequality-converter/testing/file"
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertMultipleFiles(t *testing.T) {
	files := file.NewFakeIO()
	target := NewCommand(files)

	var _ = target.Convert([]string{
		"sarif-converter",
		"--type=codequality",
		"semgrep.sarif",
		"security-scan.sarif",
		"output.json",
	})
	actual, _ := files.Read("output.json")

	assert.Equal(t, fixture.MultiRunCodeQuality(), string(actual))
}
