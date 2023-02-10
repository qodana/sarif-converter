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

	assert.Equal(t, []string{"semgrep.sarif"}, arguments.Input().Paths())
}

func TestArguments_Output(t *testing.T) {
	arguments, _ := Parse([]string{"codequality-converter", "semgrep.sarif", "gl-code-quality.json"})

	assert.Equal(t, "gl-code-quality.json", arguments.Output())
}

func TestArguments_TypeDefault(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "html", arguments.Type())
}

func TestArguments_TypeSast(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--type", "sast", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "sast", arguments.Type())
}

func TestArguments_TypeHtml(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--type", "html", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "html", arguments.Type())
}

func TestArguments_TypeInvalid(t *testing.T) {
	_, err := Parse([]string{"codequaility-converter", "--type", "invalid", "semgrep.sarf", "gl-sast-report.json"})

	assert.Error(t, err)
}

func TestArguments_SrcRootRelative(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--src-root", ".", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "file:///foo/bar", *arguments.SrcRoot("/foo/bar"))
}

func TestArguments_SrcRootParent(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--src-root", "..", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "file:///foo", *arguments.SrcRoot("/foo/bar"))
}

func TestArguments_SrcRootAbsolute(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "--src-root=/home/john", "semgrep.sarf", "gl-sast-report.json"})

	assert.Equal(t, "file:///home/john", *arguments.SrcRoot("/foo/bar"))
}

func TestArguments_SrcRootNone(t *testing.T) {
	arguments, _ := Parse([]string{"codequaility-converter", "semgrep.sarf", "gl-sast-report.json"})

	assert.Nil(t, arguments.SrcRoot("/foo/bar"))
}
