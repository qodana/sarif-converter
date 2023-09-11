package converter

import (
	"sarif-converter/meta"
	"sarif-converter/now"
	"sarif-converter/sarifreport/report"
	"sarif-converter/sast"
)

type sastConverter struct {
	time     *now.TimeProvider
	metadata meta.Metadata
}

var Sast = sastConverter{}

func (c sastConverter) Type() string {
	return "sast"
}

func (c sastConverter) Convert(input []byte) ([]byte, error) {
	sarifReport, err := report.FromBytes(input)
	if err != nil {
		return nil, err
	}

	r, err := sast.ConvertFrom(sarifReport, c.time, c.metadata)
	if err != nil {
		return nil, err
	}

	return r.Json()
}

func (c sastConverter) WithTimeProvider(p *now.TimeProvider) Converter {
	if p != nil {
		c.time = p
	}
	return c
}

func (c sastConverter) WithMetadata(metadata *meta.Metadata) sastConverter {
	if metadata != nil {
		c.metadata = *metadata
	}
	return c
}
