package run

import "sarif-converter/sarifreport/result"

func messages(results result.Wrappers) []string {
	list := make([]string, 0)

	for wrapper := range results.Iter() {
		list = append(list, *wrapper.TextMessage())
	}

	return list
}

func p(s string) *string {
	return &s
}
