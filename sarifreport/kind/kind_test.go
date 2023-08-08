package kind

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKind(t *testing.T) {
	actual := GetKind(&sarif.Result{Kind: p("pass")})

	assert.Equal(t, "pass", actual)
}

func TestGetKind_KindIsAbsent(t *testing.T) {
	actual := GetKind(&sarif.Result{})

	assert.Equal(t, "fail", actual)
}

func p(s string) *string {
	return &s
}
