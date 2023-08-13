package codequality

import (
	"codequality-converter/sarifreport/result"
)

type issue struct {
	r result.Wrapper
}

func (i issue) element() Element {
	element := Element{
		Description: i.description(),
		Severity:    i.severity(),
		Location:    i.location(),
	}

	element.Fingerprint = Fingerprint(element)

	return element
}

func (i issue) description() *string {
	message := i.r.TextMessage()
	if message != nil {
		return message
	}

	return i.r.Rule().TextFullDescription()
}

func (i issue) severity() string {
	return severity(i.r.Level())
}

func (i issue) location() Location {
	location := i.r.FirstLocation()
	return Location{
		Path:  location.PhysicalLocation.ArtifactLocation.URI,
		Lines: i.locationLine(),
	}
}

func (i issue) locationLine() *LocationLine {
	line := i.line()
	if line == nil {
		return nil
	}
	return &LocationLine{
		Begin: *line,
	}
}

func (i issue) line() *int {
	location := i.r.FirstLocation()
	if location.PhysicalLocation == nil {
		return nil
	}
	if location.PhysicalLocation.Region == nil {
		return nil
	}

	return location.PhysicalLocation.Region.StartLine
}

func newIssue(r result.Wrapper) issue {
	return issue{r: r}
}
