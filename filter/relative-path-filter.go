package filter

import (
	"codequality-converter/sarifreport/originaluri"
	"github.com/owenrumney/go-sarif/v2/sarif"
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
		bases := originaluri.NewBases(run.OriginalUriBaseIDs)
		for _, result := range run.Results {
			for _, location := range result.Locations {
				l := location.PhysicalLocation.ArtifactLocation
				f.convertToRelativePath2(l, bases)
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

func (f *RelativePathFilter) convertToRelativePath2(location *sarif.ArtifactLocation, bases originaluri.Bases) {
	fullPath := newArtifactLocationWrapper(location).resolve(bases)

	relative, _ := filepath.Rel(f.srcRoot, fullPath)
	slashed := filepath.ToSlash(relative)

	sr := "%SRCROOT%"

	location.URI = &slashed
	location.URIBaseId = &sr
}
