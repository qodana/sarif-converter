package issue

import (
	"sarif-converter/codequality/element"
	"sarif-converter/codequality/fingerprint"
	"sarif-converter/codequality/severity"
	"sarif-converter/sarifreport/result"
)

type Issue struct {
	r result.Wrapper
}

func (i Issue) Element() element.Element {
	el := element.Element{
		CheckName:   i.checkName(),
		Description: i.description(),
		Severity:    i.severity(),
		Location:    i.location(),
	}

	el.Fingerprint = fingerprint.Fingerprint(el)

	return el
}

func (i Issue) description() *string {
	message := i.r.TextMessage()
	if message != nil {
		return message
	}

	return i.r.Rule().TextFullDescription()
}

func (i Issue) severity() string {
	return severity.GetSeverity(i.r.Level())
}

func (i Issue) location() element.Location {
	location := i.r.FirstLocation()
	return element.Location{
		Path:  location.PhysicalLocation.ArtifactLocation.URI,
		Lines: i.locationLine(),
	}
}

func (i Issue) checkName() *string {
	return i.r.RuleId()
}

func (i Issue) locationLine() *element.LocationLine {
	line := i.line()
	if line == nil {
		return nil
	}
	return &element.LocationLine{
		Begin: *line,
	}
}

func (i Issue) line() *int {
	location := i.r.FirstLocation()
	if location.PhysicalLocation == nil {
		return nil
	}
	if location.PhysicalLocation.Region == nil {
		return nil
	}

	return location.PhysicalLocation.Region.StartLine
}

func NewIssue(r result.Wrapper) Issue {
	return Issue{r: r}
}

func NewIssues(results result.Wrappers) []Issue {
	list := make([]Issue, 0)

	for r := range results.OnlyRequireReport().Iter() {
		list = append(list, NewIssue(r))
	}

	return list
}
