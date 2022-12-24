package converter

import "errors"

func GetConverter(converterType string) Converter {
	switch converterType {
	case CodeQuality.Type():
		return CodeQuality
	case Sast.Type():
		return Sast
	case Html.Type():
		return Html
	}

	panic(errors.New("unknown converter type (" + converterType + ")"))
}
