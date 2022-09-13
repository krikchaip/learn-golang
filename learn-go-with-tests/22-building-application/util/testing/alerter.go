package testing

import (
	"fmt"
	"time"
)

type SpyBlindAlerter struct {
	Alerts []ScheduleAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduleAlert{duration, amount})
}

// implements: Stringer
type ScheduleAlert struct {
	Duration time.Duration // will forward this field to timer.AfterFunc()
	Amount   int
}

func (s ScheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.Duration)
}
