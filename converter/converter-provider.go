package converter

import (
	"errors"
	"sarif-converter/now"
)

func GetConverter(converterType string, p *now.TimeProvider) Converter {
	switch converterType {
	case CodeQuality.Type():
		return CodeQuality
	case Sast.Type():
		return Sast.WithTimeProvider(p)
	case Html.Type():
		return Html
	}

	panic(errors.New("unknown converter type (" + converterType + ")"))
}
