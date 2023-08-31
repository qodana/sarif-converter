package converter

import (
	"errors"
	"sarif-converter/meta"
	"sarif-converter/now"
)

func GetConverter(converterType string, p *now.TimeProvider, metadata *meta.Metadata) Converter {
	switch converterType {
	case CodeQuality.Type():
		return CodeQuality
	case Sast.Type():
		return Sast.WithMetadata(metadata).WithTimeProvider(p)
	case Html.Type():
		return Html
	}

	panic(errors.New("unknown converter type (" + converterType + ")"))
}
