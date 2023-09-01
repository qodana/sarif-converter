package scanning

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/now"
	"time"
)

type Scanning struct {
	timeProvider now.TimeProvider
	report       *sarif.Report
}

func (s Scanning) Status() string {
	return "success"
}

func (s Scanning) StartTime() *report.ScanTime {
	return s.scanTime(s.currentInvocation().StartTimeUTC)
}

func (s Scanning) EndTime() *report.ScanTime {
	return s.scanTime(s.currentInvocation().EndTimeUTC)
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

func (s Scanning) currentInvocation() *sarif.Invocation {
	run := s.currentRun()

	if len(run.Invocations) <= 0 {
		return emptyInvocation()
	}
	return run.Invocations[0]
}

func (s Scanning) currentRun() *sarif.Run {
	if len(s.report.Runs) <= 0 {
		return &sarif.Run{Invocations: []*sarif.Invocation{}}
	}
	return s.report.Runs[0]
}

func NewScanning(r *sarif.Report) Scanning {
	return Scanning{report: r}
}

func emptyInvocation() *sarif.Invocation {
	return &sarif.Invocation{}
}
