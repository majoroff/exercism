package clock

import "fmt"

// Clock structure
type Clock struct {
	h int
	m int
}

// New creates new Clock
func New(h, m int) Clock {

	// minutes routine
	var rolledOverHours int
	if m > 60 || m < -60 {
		rolledOverHours += m / 60
		m = m % 60
	} else if m == 60 {
		m = 0
		rolledOverHours++
	} else if m == -60 {
		rolledOverHours--
		m = 0
	}
	if m < 0 {
		rolledOverHours--
		m = 60 + m
	}

	// hours routine
	h = h + rolledOverHours
	if h > 24 || h < -24 {
		h = h % 24
	} else if h == 24 || h == -24 {
		h = 0
	}
	if h < 0 {
		h = 24 + h
	}

	return Clock{h: h, m: m}
}

// Add adds minutes to Clock
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

// Subtract subtracts minutes from Clock
func (c Clock) Subtract(m int) Clock {
	if m > 0 {
		return New(c.h, c.m-m)
	}
	return New(c.h, c.m+m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
