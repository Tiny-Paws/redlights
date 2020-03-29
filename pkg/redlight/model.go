package redlight

import "time"

// Defining color consts
const (
	Red = 1 << iota
	Orange
	Green
)

// Redlight is the redlight object duh
type Redlight struct {
	pattern     ColorPattern
	initialTime time.Time
}

// ColorPattern creates a pattern that has to be impremented by aRedlight
// [][]int must be of the form [color, duration]
type ColorPattern [][]int
