package cli

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter schedules alerts for blind amounts.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter.
func (f BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	f(duration, amount)
}

// StdOutAlerter prints the blind amount to standard output with a given duration.
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		_, err := fmt.Fprintf(os.Stdout, "Blind amount: %d\n", amount)
		if err != nil {
			panic(err)
		}
	})
}
