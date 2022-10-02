package entity

import (
	"fmt"
	"io"
	"time"
)

// ?? Remember that any type can implement an interface, not just structs.
// ?? If you are making a library that exposes an interface with one function defined
// ?? it is a common idiom to also expose a MyInterfaceFunc type.

type BlindAlerter interface {
	ScheduleAlertAt(to io.Writer, duration time.Duration, amount int)
}

// ?? This type will be a func which will also implement your interface.
// ?? That way users of your interface have the option to implement your interface
// ?? with just a function; rather than having to create an empty struct type.

// implements: entity.BlindAlerter
type BlindAlerterFunc func(to io.Writer, duration time.Duration, amount int)

func (f BlindAlerterFunc) ScheduleAlertAt(to io.Writer, duration time.Duration, amount int) {
	f(to, duration, amount)
}

var Alerter = BlindAlerterFunc(func(to io.Writer, duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
})
