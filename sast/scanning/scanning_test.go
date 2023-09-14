package scanning

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"sarif-converter/now"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	target := NewScanning(fakeReport(&sarif.Invocation{}))

	assert.Equal(t, "success", target.Status())
}

func TestStartTime(t *testing.T) {
	s := parse("2016-02-08T16:08:25.943Z")

	target := NewScanning(fakeReport(&sarif.Invocation{StartTimeUTC: &s}))

	assert.Equal(t, (report.ScanTime)(s), *target.StartTime())
}

func TestStartTime_NoTime(t *testing.T) {
	s := parse("2016-10-08T16:08:25.943Z")
	fakeTime := now.NewFakeTime(s)

	target := NewScanning(fakeReport(&sarif.Invocation{})).WithTimeProvider(fakeTime)

	assert.Equal(t, (report.ScanTime)(s), *target.StartTime())
}

func TestEndTime(t *testing.T) {
	s := parse("2016-02-08T16:08:25.943Z")

	target := NewScanning(fakeReport(&sarif.Invocation{EndTimeUTC: &s}))

	assert.Equal(t, (report.ScanTime)(s), *target.EndTime())
}

func TestEndTime_NoTime(t *testing.T) {
	s := parse("2016-10-08T16:08:25.943Z")
	fakeTime := now.NewFakeTime(s)

	target := NewScanning(fakeReport(&sarif.Invocation{})).WithTimeProvider(fakeTime)

	assert.Equal(t, (report.ScanTime)(s), *target.EndTime())
}

func TestEndTime_EmptyRun(t *testing.T) {
	s := parse("2016-10-08T16:08:25.943Z")
	fakeTime := now.NewFakeTime(s)

	target := NewScanning(&sarif.Report{Runs: []*sarif.Run{}}).WithTimeProvider(fakeTime)

	assert.Equal(t, (report.ScanTime)(s), *target.EndTime())
}

func parse(value string) time.Time {
	s, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return s
}

func fakeReport(invocation *sarif.Invocation) *sarif.Report {
	invocations := make([]*sarif.Invocation, 0)
	if invocations != nil {
		invocations = append(invocations, invocation)
	}
	return &sarif.Report{Runs: []*sarif.Run{
		{
			Invocations: invocations,
		},
	}}
}
