package testing

import (
	"fmt"
	"io"
	"time"
)

// implements: entity.BlindAlerter
type SpyBlindAlerter struct {
	Alerts []ScheduleAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(to io.Writer, duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduleAlert{duration, amount})
}

// implements: Stringer
type ScheduleAlert struct {
	At     time.Duration // will forward this field to timer.AfterFunc()
	Amount int
}

func (s ScheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}
