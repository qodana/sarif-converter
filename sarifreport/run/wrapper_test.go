package run

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrappreResults(t *testing.T) {
	target := newWrapper(&sarif.Run{
		Results: []*sarif.Result{
			{Message: sarif.Message{Text: p("result1")}},
			{Message: sarif.Message{Text: p("result2")}},
		},
	})

	assert.Equal(t, []string{"result1", "result2"}, messages(target.Results))
}
