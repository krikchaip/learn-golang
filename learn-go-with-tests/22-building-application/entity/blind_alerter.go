package entity

import "time"

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
