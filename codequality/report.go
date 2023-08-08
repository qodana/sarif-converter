package codequality

import (
	"codequality-converter/sarifreport/report"
	"encoding/json"
)

type Report struct {
	issues []issue
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

func ConvertFrom(report *report.Wrapper) Report {
	list := make([]issue, 0)

	for result := range report.Results().Iter() {
		list = append(list, newIssueWrapper(result))
	}

	return Report{issues: list}
}
