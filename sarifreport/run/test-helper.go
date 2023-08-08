package run

import "codequality-converter/sarifreport/result"

func messages(results result.Wrappers) []string {
	list := make([]string, 0)

	for wrapper := range results.Iter() {
		list = append(list, *wrapper.Message())
	}

	return list
}

func p(s string) *string {
	return &s
}
