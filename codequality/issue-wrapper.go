package codequality

import "codequality-converter/sarifreport"

type issueWrapper struct {
	issue sarifreport.Issue
}

func (i issueWrapper) element() Element {
	element := Element{
		Description: i.description(),
		Severity:    i.severity(),
		Location:    i.location(),
	}

	element.Fingerprint = Fingerprint(element)

	return element
}

func (i issueWrapper) description() *string {
	return i.issue.Message()
}

func (i issueWrapper) severity() string {
	return severity(i.issue.Level())
}

func (i issueWrapper) location() Location {
	location := i.issue.Location()
	return Location{
		Path: location.Path,
		Lines: LocationLine{
			Begin: location.StartLine,
		},
	}
}

func newIssueWrapper(issue sarifreport.Issue) issueWrapper {
	return issueWrapper{issue: issue}
}
