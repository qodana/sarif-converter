package severity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeverityError(t *testing.T) {
	assert.Equal(t, "critical", GetSeverity("error"))
}

func TestSeverityWarning(t *testing.T) {
	assert.Equal(t, "major", GetSeverity("warning"))
}

func TestSeverityNote(t *testing.T) {
	assert.Equal(t, "minor", GetSeverity("note"))
}

func TestSeverityNone(t *testing.T) {
	assert.Equal(t, "info", GetSeverity("none"))
}
