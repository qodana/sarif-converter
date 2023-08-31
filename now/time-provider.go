package now

import "time"

type TimeProvider interface {
	UtcNow() time.Time
}

func NewTimeProvider() TimeProvider {
	return timeProviderImpl{}
}
