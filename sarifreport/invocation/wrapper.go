package invocation

import "github.com/owenrumney/go-sarif/v2/sarif"

type Wrapper struct {
	invocation *sarif.Invocation
}

func (w Wrapper) exists() bool {
	return w.invocation != nil
}

func (w Wrapper) FindConfiguration(descriptor *sarif.ReportingDescriptor) *sarif.ReportingConfiguration {
	if descriptor == nil {
		return nil
	}

	for _, override := range w.invocation.RuleConfigurationOverrides {
		if *override.Descriptor.Id == descriptor.ID {
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
