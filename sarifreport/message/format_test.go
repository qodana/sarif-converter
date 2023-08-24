package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat_Simple(t *testing.T) {
	actual := format("simple", []string{})

	assert.Equal(t, "simple", actual)
}

func TestFormat_Replace(t *testing.T) {
	actual := format("{0} is {1} ({0})", []string{"a", "b"})

	assert.Equal(t, "a is b (a)", actual)
}

func TestFormat_Brackets(t *testing.T) {
	actual := format("{{}} {{", []string{})

	assert.Equal(t, "{} {", actual)
}

func TestFormat_NoNesting(t *testing.T) {
	actual := format("{0}", []string{"{1}", "A"})

	assert.Equal(t, "{1}", actual)
}

func TestFormat_Overflow_Arguments(t *testing.T) {
	actual := format("{0}", []string{})

	assert.Equal(t, "{0}", actual)
}

func TestFormat_Invalid(t *testing.T) {
	actual := format("{ 0}", []string{"a"})

	assert.Equal(t, "{ 0}", actual)
}
