package filter

import (
	"github.com/owenrumney/go-sarif/sarif"
	"path/filepath"
)

type RelativePathFilter struct {
	srcRoot string
}

func NewRelativePathFilter(srcRoot string) RelativePathFilter {
	return RelativePathFilter{srcRoot: srcRoot}
}

func (f *RelativePathFilter) Run(report *sarif.Report) {
	for _, run := range report.Runs {
		for _, result := range run.Results {
			for _, location := range result.Locations {
				l := location.PhysicalLocation.ArtifactLocation
				f.convertToRelativePath(l)
			}
		}
	}
}

func (f *RelativePathFilter) convertToRelativePath(location *sarif.ArtifactLocation) {
	relative, _ := filepath.Rel(f.srcRoot, *location.URI)
	slashed := filepath.ToSlash(relative)

	sr := "%SRCROOT%"

	location.URI = &slashed
	location.URIBaseId = &sr
}
