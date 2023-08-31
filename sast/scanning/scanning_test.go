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
	target := NewScanning(&sarif.Invocation{})

	assert.Equal(t, "success", target.Status())
}

func TestStartTime(t *testing.T) {
	s := parse("2016-02-08T16:08:25.943Z")

	target := NewScanning(&sarif.Invocation{StartTimeUTC: &s})

	assert.Equal(t, (report.ScanTime)(s), *target.StartTime())
}

func TestStartTime_NoTime(t *testing.T) {
	s := parse("2016-10-08T16:08:25.943Z")

	target := NewScanningWithTimeProvider(&sarif.Invocation{}, now.NewFakeTime(s))

	assert.Equal(t, (report.ScanTime)(s), *target.StartTime())
}

func TestEndTime(t *testing.T) {
	s := parse("2016-02-08T16:08:25.943Z")

	target := NewScanning(&sarif.Invocation{EndTimeUTC: &s})

	assert.Equal(t, (report.ScanTime)(s), *target.EndTime())
}

func TestEndTime_NoTime(t *testing.T) {
	s := parse("2016-10-08T16:08:25.943Z")

	target := NewScanningWithTimeProvider(&sarif.Invocation{}, now.NewFakeTime(s))

	assert.Equal(t, (report.ScanTime)(s), *target.EndTime())
}

func parse(value string) time.Time {
	s, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return s
}
