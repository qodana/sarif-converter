package main

import (
	"codequality-converter/testing/file"
	"codequality-converter/testing/fixture"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertMultipleFiles(t *testing.T) {
	reader := newFakeReader()
	writer := file.NewFakeWriter()
	target := newConverter(reader, writer)

	var _ = target.convert([]string{
		"sarif-converter",
		"--type=codequality",
		"semgrep.sarif",
		"security-scan.sarif",
		"output.json",
	})
	actual, _ := writer.Get("output.json")

	assert.Equal(t, fixture.MultiRunCodeQuality(), string(actual))
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
