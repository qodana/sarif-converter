package converter

import (
	"codequality-converter/testing/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertMultipleRuns(t *testing.T) {
	report, _ := GetConverter("codequality").Convert(fixture.MultiRunSarif())

	assert.Equal(t, fixture.MultiRunCodeQuality(), string(report))
}
