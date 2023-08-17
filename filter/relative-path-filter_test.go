package filter

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"sarif-converter/sarifreport/originaluri"
	"sarif-converter/testing/fixture"
	"testing"
)

func TestConvertToRelativePathArtifactLocation(t *testing.T) {
	target := NewRelativePathFilter("file:///home/masakura/tmp/sc")
	s, _ := sarif.FromBytes(fixture.SecurityCodeScan())

	target.Run(s)

	al := s.Runs[0].Results[0].Locations[0].PhysicalLocation.ArtifactLocation

	assert.Equal(t, artifactLocation{
		URI:       "Controllers/HomeController.cs",
		URIBaseID: "%SRCROOT%",
	}, extract(al))
}

func TestConvertToRelativePath(t *testing.T) {
	target := NewRelativePathFilter("file:///home/masakura/tmp/sc")

	l := sarif.ArtifactLocation{
		URI: p("file:///home/masakura/tmp/sc/Controllers/HomeController.cs"),
	}

	target.convertToRelativePath2(&l, originaluri.NewBases(map[string]*sarif.ArtifactLocation{}))

	assert.Equal(t, artifactLocation{
		URI:       "Controllers/HomeController.cs",
		URIBaseID: "%SRCROOT%",
	}, extract(&l))
}

func TestConvertToRelativePathWIthBaseURI(t *testing.T) {
	target := NewRelativePathFilter("file:///builds/jetbrains-ide-plugins/semgrep-plugin")
	bases := originaluri.NewBases(map[string]*sarif.ArtifactLocation{
		"SRCROOT": {URI: p("file:///root")},
	})

	l := sarif.ArtifactLocation{
		URI:       p("../builds/jetbrains-ide-plugins/semgrep-plugin/settings.gradle.kts"),
		URIBaseId: p("%SRCROOT%"),
	}

	target.convertToRelativePath2(&l, bases)

	assert.Equal(t, artifactLocation{
		URI:       "settings.gradle.kts",
		URIBaseID: "%SRCROOT%",
	}, extract(&l))
}

func p(s string) *string {
	return &s
}

type artifactLocation struct {
	URI       string
	URIBaseID string
}

func extract(l *sarif.ArtifactLocation) artifactLocation {
	return artifactLocation{
		URI:       *l.URI,
		URIBaseID: *l.URIBaseId,
	}
}
