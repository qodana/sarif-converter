package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func read(filename string) string {
	file, _ := os.ReadFile(filename)
	return string(file)
}

func TestConvert(t *testing.T) {
	actual := convert(read("fixtures/semgrep.sarif"))

	expected := read("fixtures/actual.json")

	assert.Equal(t, expected, actual)
}
