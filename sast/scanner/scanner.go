package scanner

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
)

type Scanner struct {
	driver *sarif.ToolComponent
}

func (s Scanner) ToSast() report.ScannerDetails {
	return report.ScannerDetails{
		ID:      s.Id(),
		Name:    s.Name(),
		Version: s.Version(),
		Vendor: report.Vendor{
			Name: s.VendorName(),
		},
	}
}

func (s Scanner) Name() string {
	return s.driver.Name
}

func (s Scanner) Version() string {
	driver := s.driver
	if driver.SemanticVersion != nil {
		return *driver.SemanticVersion
	}
	return *driver.Version
}

func (s Scanner) VendorName() string {
	driver := s.driver

	if driver.Organization != nil {
		return *driver.Organization
	}

	return driver.Name
}

func (s Scanner) Id() string {
	return s.driver.Name
}

func NewScanner(driver *sarif.ToolComponent) Scanner {
	return Scanner{driver: driver}
}

func NewScannerFrom(r *sarif.Report) Scanner {
	return NewScanner(r.Runs[0].Tool.Driver)
}
