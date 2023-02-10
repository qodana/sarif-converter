package originaluri

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolve(t *testing.T) {
	target := newBase(&sarif.ArtifactLocation{
		URI: p("file:///home/user/"),
	})

	assert.Equal(t, "file:///home/user/project1/file.txt", target.Resolve("project1/file.txt"))
}

func TestResolveNil(t *testing.T) {
	target := newBase(&sarif.ArtifactLocation{})

	assert.Equal(t, "project1/file.txt", target.Resolve("project1/file.txt"))
}

func TestResolveWithParentDirectory(t *testing.T) {
	target := newBase(&sarif.ArtifactLocation{
		URI: p("file:///root/"),
	})

	assert.Equal(t, "file:///builds/project1/file.txt", target.Resolve("../builds/project1/file.txt"))
}
