package now

import "time"

type timeProviderImpl struct {
}

func (p timeProviderImpl) UtcNow() time.Time {
	return time.Now()
}
