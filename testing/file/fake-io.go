package file

import (
	"codequality-converter/testing/fixture"
	"errors"
)

type FakeIO struct {
	files map[string][]byte
}

func (io FakeIO) Read(name string) ([]byte, error) {
	bytes, ok := io.files[name]
	if ok {
		return bytes, nil
	}
	return nil, errors.New("`" + name + "` is not found")
}

func (io FakeIO) Write(name string, data []byte) error {
	io.files[name] = data
	return nil
}

func NewFakeIO() FakeIO {
	return FakeIO{
		files: map[string][]byte{
			"semgrep.sarif":       fixture.SemgrepSarif(),
			"security-scan.sarif": fixture.SecurityCodeScan(),
			"multi-run.sarif":     fixture.MultiRunSarif(),
		},
	}
}
