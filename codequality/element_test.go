package codequality

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoLines(t *testing.T) {
	element := Element{
		Description: p("description"),
		Fingerprint: "fingerprint",
		Severity:    "info",
		Location: Location{
			Path: p("foo.js"),
		},
	}

	bytes, _ := json.Marshal(element)

	assert.Equal(t, "{\"description\":\"description\",\"fingerprint\":\"fingerprint\",\"severity\":\"info\",\"location\":{\"path\":\"foo.js\"}}", string(bytes))
}
