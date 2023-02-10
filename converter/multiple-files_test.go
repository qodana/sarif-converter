package converter

import (
	"codequality-converter/file"
	"codequality-converter/testing/fixture"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertMultipleFiles(t *testing.T) {
	reader := newFakeReader()
	inputs := file.NewInput([]string{"semgrep.sarif", "security-scan.sarif"})
	bytes, _ := inputs.Read(reader)
	report, _ := GetConverter("codequality").Convert(bytes)

	assert.Equal(t, fixture.MultiRunCodeQuality(), string(report))
}

type fakeReader struct {
}

// TODO go to testing/fixture
func (r fakeReader) Read(name string) ([]byte, error) {
	if name == "semgrep.sarif" {
		return fixture.SemgrepSarif(), nil
	}
	if name == "security-scan.sarif" {
		return fixture.SecurityCodeScan(), nil
	}
	return nil, errors.New("`" + name + "` is not found.")
}

func newFakeReader() fakeReader {
	return fakeReader{}
}
