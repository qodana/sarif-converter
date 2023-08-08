package run

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrappersResults(t *testing.T) {
	target := NewWrappers(&sarif.Report{
		Runs: []*sarif.Run{
			{
				Results: []*sarif.Result{
					{Message: sarif.Message{Text: p("run1 message1")}},
					{Message: sarif.Message{Text: p("run1 message2")}},
				},
			}, {
				Results: []*sarif.Result{
					{Message: sarif.Message{Text: p("run2 message1")}},
				},
			},
		},
	})

	assert.Equal(t, []string{"run1 message1", "run1 message2", "run2 message1"}, messages(target.Wrappers()))
}
