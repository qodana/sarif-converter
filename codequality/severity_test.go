package codequality

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeverityError(t *testing.T) {
	assert.Equal(t, "critical", severity("error"))
}

func TestSeverityWarning(t *testing.T) {
	assert.Equal(t, "major", severity("warning"))
}

func TestSeverityNote(t *testing.T) {
	assert.Equal(t, "minor", severity("note"))
}

func TestSeverityNone(t *testing.T) {
	assert.Equal(t, "info", severity("none"))
}
