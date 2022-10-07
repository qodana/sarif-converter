package argument

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArguments_TooShort(t *testing.T) {
	arguments := Parse([]string{"codequality-converter"})

	assert.False(t, arguments.IsValid())
}

func TestArguments_RequireShowVersion(t *testing.T) {
	arguments := Parse([]string{"codequality-converter", "--version"})

	assert.True(t, arguments.RequireShowVersion())
}

func TestArguments_RequireShowVersionShort(t *testing.T) {
	arguments := Parse([]string{"codequality-converter", "-v"})

	assert.True(t, arguments.RequireShowVersion())
}

func TestArguments_Input(t *testing.T) {
	arguments := Parse([]string{"codequality-converter", "semgrep.sarif", "gl-code-quality.json"})

	assert.Equal(t, "semgrep.sarif", arguments.Input())
}

func TestArguments_Output(t *testing.T) {
	arguments := Parse([]string{"codequality-converter", "semgrep.sarif", "gl-code-quality.json"})

	assert.Equal(t, "gl-code-quality.json", arguments.Output())
}
