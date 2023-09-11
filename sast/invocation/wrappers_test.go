package invocation

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWrappers_StartTimeUTC(t *testing.T) {

	t1 := parse("2021-10-11T10:00:12Z")

	target := NewWrappers([]*sarif.Invocation{
		{StartTimeUTC: &t1},
	})

	assert.Equal(t, t1, *target.StartTimeUTC())
}

func TestWrappers_StartTimeUTC_Empty(t *testing.T) {
	target := NewWrappers([]*sarif.Invocation{})

	assert.Nil(t, target.StartTimeUTC())
}

func TestWrappers_StartTimeUTC_Minimal(t *testing.T) {
	t1 := parse("2021-12-22T10:10:11Z")
	t2 := parse("2021-12-22T10:10:10Z")

	target := NewWrappers([]*sarif.Invocation{
		{StartTimeUTC: &t1},
		{StartTimeUTC: nil},
		{StartTimeUTC: &t2},
	})

	assert.Equal(t, t2, *target.StartTimeUTC())
}

func TestWrappers_EndTimeUTC(t *testing.T) {

	t1 := parse("2021-10-11T10:00:12Z")

	target := NewWrappers([]*sarif.Invocation{
		{EndTimeUTC: &t1},
	})

	assert.Equal(t, t1, *target.EndTimeUTC())
}

func TestWrappers_EndTimeUTC_Empty(t *testing.T) {
	target := NewWrappers([]*sarif.Invocation{})

	assert.Nil(t, target.EndTimeUTC())
}

func TestWrappers_EndTimeUTC_Minimal(t *testing.T) {
	t1 := parse("2021-12-22T10:10:09Z")
	t2 := parse("2021-12-22T10:10:10Z")

	target := NewWrappers([]*sarif.Invocation{
		{EndTimeUTC: &t1},
		{EndTimeUTC: nil},
		{EndTimeUTC: &t2},
	})

	assert.Equal(t, t2, *target.EndTimeUTC())
}

func parse(value string) time.Time {
	s, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return s
}
