package sarifreport

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoResults(t *testing.T) {
	report := NewReport(&sarif.Report{
		Runs: []*sarif.Run{
			{},
		},
	})

	assert.NotNil(t, report.CodeQualityElements())
}

func TestZeroResults(t *testing.T) {
	report := NewReport(&sarif.Report{
		Runs: []*sarif.Run{
			{
				Results: []*sarif.Result{},
			},
		},
	})

	assert.NotNil(t, report.CodeQualityElements())
}
