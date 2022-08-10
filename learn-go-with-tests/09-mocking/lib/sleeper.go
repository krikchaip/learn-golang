package countdown

import "time"

type Sleeper interface {
	Sleep()
}

// implements: countdown.Sleeper
type SecondSleeper struct {
	Duration time.Duration
}

func (ss SecondSleeper) Sleep() {
	time.Sleep(ss.Duration * time.Second)
}

// implements: countdown.Sleeper
type ConfigurableSleeper struct {
	Duration time.Duration
	SleepFn  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFn(c.Duration)
}
