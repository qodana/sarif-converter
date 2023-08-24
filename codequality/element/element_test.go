package element

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoLines(t *testing.T) {
	element := Element{
		CheckName:   p("security/detect-eval-with-expression"),
		Description: p("description"),
		Fingerprint: "fingerprint",
		Severity:    "info",
		Location: Location{
			Path: p("foo.js"),
		},
	}

	bytes, _ := json.Marshal(element)

	assert.Equal(t, "{\"check_name\":\"security/detect-eval-with-expression\",\"description\":\"description\",\"fingerprint\":\"fingerprint\",\"severity\":\"info\",\"location\":{\"path\":\"foo.js\"}}", string(bytes))
}

func p(s string) *string {
	return &s
}
