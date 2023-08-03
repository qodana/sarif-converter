package run

import (
	"codequality-converter/sarifreport/result"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Wrapper struct {
	run     *sarif.Run
	Results result.Wrappers
}
