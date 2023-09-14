package meta

import "strings"

type Metadata struct {
	Package  string
	Name     string
	Version  string
	Revision string
}

func (m Metadata) SemanticVersion() string {
	return strings.Replace(m.Version, "v", "", 1)
}

func NewMetadata(version string, revision string) Metadata {
	return Metadata{
		Package:  "sarif-converter",
		Name:     "SARIF Converter",
		Version:  version,
		Revision: revision,
	}
}
