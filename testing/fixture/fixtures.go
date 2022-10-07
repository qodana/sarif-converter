package fixture

import "os"

type Fixtures struct {
	path string
}

func (f Fixtures) ActualJson() string {
	bytes, _ := os.ReadFile(f.path + "/actual.json")
	return string(bytes)
}

func (f Fixtures) SemgrepSarif() []byte {
	bytes, _ := os.ReadFile(f.path + "/semgrep.sarfi")
	return bytes
}

func NewFixtures(path string) Fixtures {
	return Fixtures{path: path}
}
