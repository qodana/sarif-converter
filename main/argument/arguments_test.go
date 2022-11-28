package argument

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArguments_TooShort(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter"})

	assert.False(t, arguments.IsValid())
}

func TestArguments_RequireShowVersion(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter", "--version"})

	assert.True(t, arguments.RequireShowVersion())
}

func TestArguments_RequireShowVersionShort(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter", "-v"})

	assert.True(t, arguments.RequireShowVersion())
}

func TestArguments_Input(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter", "semgrep.sarif", "gl-code-quality.json"})

	assert.Equal(t, "semgrep.sarif", arguments.Input())
}

func TestArguments_Output(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter", "semgrep.sarif", "gl-code-quality.json"})

	assert.Equal(t, "gl-code-quality.json", arguments.Output())
}

func TestArguments_TypeDefault(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "codequality", arguments.Type())
}

func TestArguments_TypeSast(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--type", "sast", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "sast", arguments.Type())
}

func TestArguments_TypeInvalid(t *testing.T) {
	_, err := Parse([]string{"codequaility-converter", "--type", "invalid", "semgrep.sarf", "gl-sast-report.json"})

	assert.Error(t, err)
}
