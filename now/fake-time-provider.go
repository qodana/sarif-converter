package now

import "time"

type fakeTimeProvider struct {
	time time.Time
}

func (p fakeTimeProvider) UtcNow() time.Time {
	return p.time
}

func NewFakeTime(t time.Time) TimeProvider {
	return fakeTimeProvider{time: t}
}

func Parse(value string) time.Time {
	s, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return s
}
