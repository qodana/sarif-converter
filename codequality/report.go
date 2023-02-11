package codequality

import (
	"codequality-converter/sarifreport"
	"encoding/json"
)

type Report struct {
	issues []issueWrapper
}

func (r Report) Json() ([]byte, error) {
	return json.MarshalIndent(r.elements(), "", "  ")
}

func (r Report) elements() []Element {
	result := make([]Element, len(r.issues))

	for i, issue := range r.issues {
		result[i] = issue.element()
	}

	return result
}

func ConvertFrom(report sarifreport.ReportWrapper) Report {
	issues := report.Issues()
	result := make([]issueWrapper, len(issues))

	for i, issue := range issues {
		result[i] = newIssueWrapper(issue)
	}

	return Report{issues: result}
}
