package converter

import (
	"errors"
	"sarif-converter/converter/codequality"
	"sarif-converter/converter/html"
	"sarif-converter/converter/sast"
	"sarif-converter/meta"
	"sarif-converter/now"
)

func GetConverter(converterType string, metadata meta.Metadata) Converter {
	switch converterType {
	case codequality.CodeQuality:
		return codequality.NewCodeQualityConverter()
	case sast.Sast:
		return sast.NewSastConverter(now.NewTimeProvider(), metadata)
	case html.Html:
		return html.NewHtmlConverter()
	}

	panic(errors.New("unknown converter type (" + converterType + ")"))
}
