package invocation

import "github.com/owenrumney/go-sarif/v2/sarif"

type Wrapper struct {
	invocation *sarif.Invocation
}

func (w Wrapper) exists() bool {
	return w.invocation != nil
}

func (w Wrapper) FindConfiguration(descriptorID *string) *sarif.ReportingConfiguration {
	if descriptorID == nil {
		return nil
	}

	for _, override := range w.invocation.RuleConfigurationOverrides {
		if *override.Descriptor.Id == *descriptorID {
			return override.Configuration
		}
	}
	return nil
}

func newWrapper(invocation *sarif.Invocation) Wrapper {
	return Wrapper{
		invocation: invocation,
	}
}
