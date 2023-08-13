package rule

import "github.com/owenrumney/go-sarif/v2/sarif"

type Wrapper struct {
	rule *sarif.ReportingDescriptor
}

func (w Wrapper) name() *string {
	if w.rule == nil {
		return nil
	}

	return w.rule.Name
}

func (w Wrapper) Is(id *string) bool {
	return w.rule.ID == *id
}

func (w Wrapper) exists() bool {
	return w.rule != nil
}

func (w Wrapper) DefaultLevel() *string {
	configuration := w.defaultConfiguration()
	if configuration == nil {
		return nil
	}
	return &configuration.Level
}

func (w Wrapper) defaultConfiguration() *sarif.ReportingConfiguration {
	if w.rule == nil {
		return nil
	}
	return w.rule.DefaultConfiguration
}

func (w Wrapper) ID() *string {
	if w.rule == nil {
		return nil
	}
	return &w.rule.ID
}

func (w Wrapper) TextFullDescription() *string {
	rule := w.rule
	if rule == nil {
		return nil
	}
	description := rule.FullDescription
	if description == nil {
		return nil
	}
	return description.Text
}

func newWrapper(rule *sarif.ReportingDescriptor) Wrapper {
	return Wrapper{rule: rule}
}
