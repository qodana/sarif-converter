package kind

import "github.com/owenrumney/go-sarif/v2/sarif"

func GetKind(result *sarif.Result) string {
	if result.Kind == nil {
		return "fail"
	}

	return *result.Kind
}
