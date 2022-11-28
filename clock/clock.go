package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (c *FixedClocker) Now() time.Time {
	return time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)
}
