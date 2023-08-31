package scanning

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/now"
	"time"
)

type Scanning struct {
	invocation   *sarif.Invocation
	timeProvider now.TimeProvider
}

func (s Scanning) Status() string {
	return "success"
}

func (s Scanning) StartTime() *report.ScanTime {
	return s.scanTime(s.invocation.StartTimeUTC)
}

func (s Scanning) EndTime() *report.ScanTime {
	return s.scanTime(s.invocation.EndTimeUTC)
}

func (s Scanning) scanTime(t *time.Time) *report.ScanTime {
	if t != nil {
		return (*report.ScanTime)(t)
	}

	n := s.timeProvider.UtcNow()
	return (*report.ScanTime)(&n)
}

func (s Scanning) Override(scan report.Scan) report.Scan {
	scan.Status = report.Status(s.Status())
	scan.StartTime = s.StartTime()
	scan.EndTime = s.EndTime()
	return scan
}

func (s Scanning) WithTimeProvider(provider *now.TimeProvider) Scanning {
	s.timeProvider = *provider
	return s
}

func NewScanning(invocation *sarif.Invocation) Scanning {
	return NewScanningWithTimeProvider(invocation, now.NewTimeProvider())
}

func NewScanningFrom(r *sarif.Report) Scanning {
	return NewScanning(r.Runs[0].Invocations[0])
}

func NewScanningWithTimeProvider(invocation *sarif.Invocation, p now.TimeProvider) Scanning {
	return Scanning{
		invocation:   invocation,
		timeProvider: p,
	}
}
