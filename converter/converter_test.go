package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func read(filename string) []byte {
	file, _ := os.ReadFile(filename)
	return file
}

func TestConvert(t *testing.T) {
	report, _ := Convert(read("fixtures/semgrep.sarif"))
	actual := string(report)

	expected := string(read("fixtures/actual.json"))

	assert.Equal(t, expected, actual)
}
