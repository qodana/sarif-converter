package invocation

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindConfiguration(t *testing.T) {
	target := newWrapper(fakeInvocation())

	configuration := target.FindConfiguration(&sarif.ReportingDescriptor{ID: "error1"})

	assert.Equal(t, "error", configuration.Level)
}

func TestFindConfiguration_None(t *testing.T) {
	target := newWrapper(fakeInvocation())

	configuration := target.FindConfiguration(&sarif.ReportingDescriptor{ID: "not-found"})

	assert.Nil(t, configuration)
}

func TestFindConfiguration_NoDescriptor(t *testing.T) {
	target := newWrapper(fakeInvocation())

	configuration := target.FindConfiguration(nil)

	assert.Nil(t, configuration)

}

func fakeInvocation() *sarif.Invocation {
	return &sarif.Invocation{
		RuleConfigurationOverrides: []*sarif.ConfigurationOverride{
			newConfigurationOverride("note1", "note"),
			newConfigurationOverride("error1", "error"),
		},
	}
}

func newConfigurationOverride(descriptorId string, defaultLevel string) *sarif.ConfigurationOverride {
	return &sarif.ConfigurationOverride{
		Descriptor: &sarif.ReportingDescriptorReference{Id: &descriptorId},
		Configuration: &sarif.ReportingConfiguration{
			Level: defaultLevel,
		},
	}
}
