package scanner

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestId(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{Name: "my tool"})

	assert.Equal(t, "my tool", target.Id())
}

func TestName(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{Name: "my tool"})

	assert.Equal(t, "my tool", target.Name())
}

func TestVersion_Version(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{
		Version: p("1.2.3"),
	})

	assert.Equal(t, "1.2.3", target.Version())
}

func TestVersion_SemanticVersion(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{
		Version:         p("202301"),
		SemanticVersion: p("1.2.3"),
	})

	assert.Equal(t, "1.2.3", target.Version())
}

func TestVendorName(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{
		Organization: p("my org"),
	})

	assert.Equal(t, "my org", target.VendorName())
}

func TestVendorName_NoOrganization(t *testing.T) {
	target := NewScanner(&sarif.ToolComponent{
		Name: "my tool",
	})

	assert.Equal(t, "my tool", target.VendorName())
}

func p(s string) *string {
	return &s
}
