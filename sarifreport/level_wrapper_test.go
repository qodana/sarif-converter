package sarifreport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeverityError(t *testing.T) {
	target := LevelWrapper{level: "error"}

	assert.Equal(t, "critical", target.Severity())
}

func TestSeverityWarning(t *testing.T) {
	target := LevelWrapper{level: "warning"}

	assert.Equal(t, "major", target.Severity())
}

func TestSeverityNote(t *testing.T) {
	target := LevelWrapper{level: "note"}

	assert.Equal(t, "minor", target.Severity())
}
