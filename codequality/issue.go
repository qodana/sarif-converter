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
	return i.r.Message()
}

func (i issue) severity() string {
	return severity(i.r.Level())
}

func (i issue) location() Location {
	location := i.r.FirstLocation()
	return Location{
		Path: location.PhysicalLocation.ArtifactLocation.URI,
		Lines: LocationLine{
			Begin: *location.PhysicalLocation.Region.StartLine,
		},
	}
}

func newIssueWrapper(r result.Wrapper) issue {
	return issue{r: r}
}
