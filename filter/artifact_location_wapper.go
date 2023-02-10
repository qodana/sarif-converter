package filter

import (
	"codequality-converter/sarifreport/originaluri"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type artifactLocationWrapper struct {
	target *sarif.ArtifactLocation
}

func (w artifactLocationWrapper) resolve(bases originaluri.Bases) string {
	return bases.Resolve(w.target)
}

func newArtifactLocationWrapper(t *sarif.ArtifactLocation) artifactLocationWrapper {
	return artifactLocationWrapper{target: t}
}
