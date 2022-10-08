package fixture

import (
	_ "embed"
)

// go:embed actual.json
var actualJson string

// go:embed semgrep.sarif
var semgrepSarif []byte

type Fixtures struct {
	path string
}

func (f Fixtures) ActualJson() string {
	return actualJson
}

func (f Fixtures) SemgrepSarif() []byte {
	return semgrepSarif
}

func NewFixtures(path string) Fixtures {
	return Fixtures{path: path}
}
