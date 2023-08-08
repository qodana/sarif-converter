package rule

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultLevel(t *testing.T) {
	target := newWrapper(&sarif.ReportingDescriptor{
		DefaultConfiguration: &sarif.ReportingConfiguration{
			Level: "error",
		},
	})

	assert.Equal(t, "error", *target.DefaultLevel())
}

func TestDefaultLevel_NoDefaultConfiguration(t *testing.T) {
	target := newWrapper(&sarif.ReportingDescriptor{})

	assert.Nil(t, target.DefaultLevel())
}

func TestDefaultLevel_NoRule(t *testing.T) {
	target := newWrapper(nil)

	assert.Nil(t, target.DefaultLevel())
}
