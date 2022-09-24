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
	actual := string(Convert(read("fixtures/semgrep.sarif")))

	expected := string(read("fixtures/actual.json"))

	assert.Equal(t, expected, actual)
}
