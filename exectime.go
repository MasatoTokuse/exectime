package exectime

import "time"

func Measure(fn func()) time.Duration {
	beganAt := time.Now()
	fn()
	return time.Since(beganAt)
}
