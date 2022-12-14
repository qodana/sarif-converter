package fixture

import (
	_ "embed"
)

//go:embed actual.json
var actualJson string

//go:embed semgrep.sarif
var semgrepSarif []byte

//go:embed resharper.sarif
var resharperSarif []byte

//go:embed gl-sast-report.json
var sast []byte

//go:embed sarif-report.html
var html string

type Fixtures struct {
	path string
}

func (f Fixtures) ActualJson() string {
	return actualJson
}

func (f Fixtures) SemgrepSarif() []byte {
	return semgrepSarif
}

func ReSharperSarif() []byte {
	return resharperSarif
}

func Sast() []byte {
	return sast
}

func Html() string {
	return html
}

func NewFixtures(path string) Fixtures {
	return Fixtures{path: path}
}
