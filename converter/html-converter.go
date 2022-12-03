package converter

import (
	_ "embed"
	"encoding/base64"
	"strings"
)

//go:embed sarif-viewer.template.html
var template string

type htmlConverter struct {
}

var Html = htmlConverter{}

func (c htmlConverter) Type() string {
	return "html"
}

func (c htmlConverter) Convert(input []byte) ([]byte, error) {
	output := strings.Replace(template, "%sarif%", base64.StdEncoding.EncodeToString(input), 1)
	return []byte(output), nil
}
