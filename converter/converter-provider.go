package converter

import "errors"

func GetConverter(converterType string) Converter {
	if converterType == CodeQuality.Type() {
		return CodeQuality
	}
	if converterType == Sast.Type() {
		return Sast
	}
	if converterType == Html.Type() {
		return Html
	}

	panic(errors.New("unknown converter type (" + converterType + ")"))
}
