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
