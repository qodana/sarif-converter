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

//go:embed resharper-no-inspections.sarif
var resharperNoInspectionsSarif []byte

//go:embed security-scan.sarif
var securityCodeScanSarif []byte

//go:embed eslint.sarif
var eslint []byte

//go:embed gl-sast-report.json
var sast []byte

//go:embed multi-run.sarif
var multiRunSarif []byte

//go:embed multi-run-actual.json
var multiRunCodeQuality string

//go:embed ktlint.sarif
var ktlintSarif []byte

//go:embed ktlint-actual.json
var ktlintCodeQuality string

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

func ReSharperNoInspectionsSarif() []byte {
	return resharperNoInspectionsSarif
}

func SemgrepSarif() []byte {
	return semgrepSarif
}

func SecurityCodeScan() []byte {
	return securityCodeScanSarif
}

func Eslint() []byte {
	return eslint
}

func Sast() []byte {
	return sast
}

func MultiRunSarif() []byte {
	return multiRunSarif
}

func MultiRunCodeQuality() string {
	return multiRunCodeQuality
}

func KtlintSarif() []byte {
	return ktlintSarif
}

func KtlintCodeQuality() string {
	return ktlintCodeQuality
}

func Html() string {
	return html
}

func NewFixtures(path string) Fixtures {
	return Fixtures{path: path}
}
