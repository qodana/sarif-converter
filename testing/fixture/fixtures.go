package fixture

import (
	_ "embed"
)

//go:embed actual.json
var actualJson string

//go:embed semgrep.sarif
var semgrepSarif []byte

//go:embed gl-sast-report.json
var sast []byte

type Fixtures struct {
	path string
}

func (f Fixtures) ActualJson() string {
	return actualJson
}

func (f Fixtures) SemgrepSarif() []byte {
	return semgrepSarif
}

func Sast() []byte {
	return sast
}

func NewFixtures(path string) Fixtures {
	return Fixtures{path: path}
}
