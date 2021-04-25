package exectime

import "time"

func MeasureExecTime(fn func()) time.Duration {
	beganAt := time.Now()
	fn()
	return time.Since(beganAt)
}
