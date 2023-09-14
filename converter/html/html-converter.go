package html

import (
	_ "embed"
	"encoding/base64"
	"strings"
)

//go:embed sarif-viewer.template.html
var template string

type HtmlConverter struct {
}

var Html = "html"

func (c HtmlConverter) Convert(input []byte) ([]byte, error) {
	output := strings.Replace(template, "%sarif%", base64.StdEncoding.EncodeToString(input), 1)
	return []byte(output), nil
}

func NewHtmlConverter() HtmlConverter {
	return HtmlConverter{}
}
