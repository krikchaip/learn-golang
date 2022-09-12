package testing

import "time"

type SpyBlindAlerter struct {
	Alerts []struct {
		duration time.Duration // will forward this field to timer.AfterFunc()
		amount   int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, struct {
		duration time.Duration
		amount   int
	}{duration, amount})
}
