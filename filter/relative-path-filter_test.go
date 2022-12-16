package filter

import (
	"codequality-converter/testing/fixture"
	"github.com/owenrumney/go-sarif/sarif"
	"github.com/stretchr/testify/assert"
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
