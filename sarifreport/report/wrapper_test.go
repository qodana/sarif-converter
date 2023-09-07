package report

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapper_Value_ReportGlobal(t *testing.T) {
	target := NewReport(&sarif.Report{
		Version: "dummy",
	})

	assert.Equal(t, "dummy", target.Value().Version)
}

func TestWrapper_Value_Result(t *testing.T) {
	target := NewReport(&sarif.Report{
		Runs: []*sarif.Run{
			{
				Results: []*sarif.Result{
					fakeResult("fail", "dummy"),
				}},
		},
	})

	assert.Equal(t, []string{"dummy"}, getResultTexts(target.Value()))
}

func TestWrapper_OnlyRequireReport(t *testing.T) {
	target := NewReport(&sarif.Report{
		Runs: []*sarif.Run{
			{
				Results: []*sarif.Result{
					fakeResult("fail", "fail result"),
					fakeResult("pass", "pass result"),
					fakeResult("open", "open result"),
				},
			},
		},
	})

	assert.Equal(t, []string{"fail result", "open result"}, getResultTexts(target.OnlyRequireReport().Value()))
}

func getResultTexts(w *sarif.Report) []string {
	r := make([]string, 0)

	for _, run := range w.Runs {
		for _, result := range run.Results {
			r = append(r, *result.Message.Text)
		}
	}
	return r
}

func fakeResult(kind string, text string) *sarif.Result {
	return &sarif.Result{
		Kind:    &kind,
		Message: sarif.Message{Text: &text},
	}
}
