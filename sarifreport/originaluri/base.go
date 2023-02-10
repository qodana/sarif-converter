package originaluri

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"net/url"
)

type Base struct {
	uri *string
}

func (b Base) Resolve(relativeUri string) string {
	if b.uri == nil {
		return relativeUri
	}

	s, _ := url.JoinPath(*b.uri, relativeUri)
	return s
}

func newBase(l *sarif.ArtifactLocation) Base {
	if l == nil {
		return Base{}
	}

	return Base{
		uri: l.URI,
	}
}
