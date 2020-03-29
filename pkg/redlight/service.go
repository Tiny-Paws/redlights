package redlight

import "time"

//NewRedlight returns a Redlight
func NewRedlight(initTime time.Time) *Redlight {
	return &Redlight{
		initialTime: initTime,
	}
}

// SetPattern is the setter function for the pattern attribute
func (r *Redlight) SetPattern(pattern ColorPattern) {
	r.pattern = pattern
}

// // Pattern is a debug func
// func (r *Redlight) Pattern() ColorPattern {
// 	return r.pattern
// }
