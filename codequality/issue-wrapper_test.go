package codequality

import (
	"codequality-converter/sarifreport"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation(t *testing.T) {
	target := newFakeIssueWrapper("foo.js", newRegion(10))

	assert.Equal(t, Location{
		Path:  p("foo.js"),
		Lines: &LocationLine{Begin: 10},
	}, target.location())
}

func TestLocationNoRegion(t *testing.T) {
	target := newFakeIssueWrapper("foo.js", nil)

	assert.Equal(t, Location{
		Path:  p("foo.js"),
		Lines: nil,
	}, target.location())
}

func newRegion(startLine int) *sarif.Region {
	return &sarif.Region{StartLine: pi(startLine)}
}

func newFakeIssueWrapper(path string, region *sarif.Region) issueWrapper {
	target := newIssueWrapper(sarifreport.NewIssue(&sarif.Result{
		Locations: []*sarif.Location{
			{
				PhysicalLocation: &sarif.PhysicalLocation{
					ArtifactLocation: &sarif.ArtifactLocation{
						URI: p(path),
					},
					Region: region,
				},
			},
		},
	}, sarifreport.SarifRunWrapper{}))
	return target
}
