package codequality

import (
	"encoding/json"
	"sarif-converter/codequality/element"
	"sarif-converter/codequality/issue"
	"sarif-converter/sarifreport/report"
)

type Report struct {
	issues []issue.Issue
}

func (r Report) Json() ([]byte, error) {
	return json.MarshalIndent(r.elements(), "", "  ")
}

func (r Report) elements() []element.Element {
	result := make([]element.Element, len(r.issues))

	for i, is := range r.issues {
		result[i] = is.Element()
	}

	return result
}

func ConvertFrom(report *report.Wrapper) Report {
	list := make([]issue.Issue, 0)

	for result := range report.Results().Iter() {
		list = append(list, issue.NewIssue(result))
	}

	return Report{issues: list}
}
