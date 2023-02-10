package command

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertKtlint(t *testing.T) {
	target := newWrapper()

	target.convert([]string{
		"sarif-converter",
		"--type=codequality",
		"--src-root=/builds/jetbrains-ide-plugins/semgrep-plugin",
		"ktlint.sarif",
		"output.json",
	})

	assert.Equal(t, fixture.KtlintCodeQuality(), target.readText("output.json"))
}
