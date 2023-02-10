package originaluri

import "github.com/owenrumney/go-sarif/v2/sarif"

type Bases struct {
	target map[string]*sarif.ArtifactLocation
}

func (b Bases) Find(key *string) Base {
	for k, l := range b.target {
		if b.match(key, k) {
			return newBase(l)
		}
	}
	return Base{}
}

func (b Bases) Resolve(l *sarif.ArtifactLocation) string {
	return b.Find(l.URIBaseId).Resolve(*l.URI)
}

func (b Bases) match(findKey *string, originalKey string) bool {
	if findKey == nil {
		return false
	}
	return originalKey == *findKey || ("%"+originalKey+"%" == *findKey)
}

func NewBases(l map[string]*sarif.ArtifactLocation) Bases {
	return Bases{target: l}
}
