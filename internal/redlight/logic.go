package redlight

import (
	"time"
)

// ColorState returns the state of the redlight at the given t time
func ColorState(r *Redlight, t time.Time) int {
	totalPatternDuration := func() (result int) {
		for _, v := range r.pattern {
			result += v[1]
		}
		return
	}()
	runtime := int(t.Sub(r.initialTime).Seconds())
	timeInLoop := runtime % totalPatternDuration

	var cumulatedTime int
	var state int
	for _, v := range r.pattern {
		if cumulatedTime > timeInLoop {
			break
		}
		cumulatedTime += v[1]
		state = v[0]
	}
	return state
}
