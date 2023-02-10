package originaluri

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	target := NewBases(map[string]*sarif.ArtifactLocation{
		"USER":    {URI: p("file:///home/user/")},
		"SRCROOT": {URI: p("file:///root/")},
	})

	r := target.Find(p("SRCROOT"))

	assert.Equal(t, p("file:///root/"), r.uri)
}

func TestFindWithPercent(t *testing.T) {
	target := NewBases(map[string]*sarif.ArtifactLocation{
		"USER":    {URI: p("file:///home/user/")},
		"SRCROOT": {URI: p("file:///root/")},
	})

	r := target.Find(p("%SRCROOT%"))

	assert.Equal(t, p("file:///root/"), r.uri)
}

func TestFindNotFound(t *testing.T) {
	target := NewBases(map[string]*sarif.ArtifactLocation{
		"USER":    {URI: p("file:///home/user/")},
		"SRCROOT": {URI: p("file:///root/")},
	})

	r := target.Find(p("abc"))

	assert.Nil(t, r.uri)
}

func p(s string) *string {
	return &s
}
