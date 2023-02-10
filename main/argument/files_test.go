package argument

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFileArguments(t *testing.T) {
	actual := parseFileArguments([]string{"input1.sarif", "input2.sarif", "output.json"})

	assert.Equal(t, files{
		inputs: []string{"input1.sarif", "input2.sarif"},
		output: "output.json",
	}, actual)
}

func TestParseFileArguments_One(t *testing.T) {
	actual := parseFileArguments([]string{"output.json"})

	assert.Equal(t, files{
		inputs: []string{},
		output: "output.json",
	}, actual)
}

func TestParseFileArguments_Empty(t *testing.T) {
	actual := parseFileArguments([]string{})

	assert.Equal(t, files{
		inputs: []string{},
		output: "",
	}, actual)
}
