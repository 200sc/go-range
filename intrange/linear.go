package intrange

import "math/rand"

// NewLinear returns a linear range
// between min and max
func NewLinear(min, max int) Range {
	if max == min {
		return Constant(min)
	}
	if max < min {
		max, min = min, max
	}
	return linear{min, max}
}

// NewSpread returns a linear range from base - s to base + s
func NewSpread(base, s int) Range {
	if s == 0 {
		return Constant(base)
	}
	if s < 0 {
		s *= -1
	}
	return linear{base - s, base + s}
}

// Linear polls on a linear scale
// between a minimum and a maximum
// linear is private so that the maximum cannot
// be changed to be less than the minimum.
type linear struct {
	Min, Max int
}

func (lir linear) Poll() int {
	return rand.Intn((lir.Max+1)-lir.Min) + lir.Min
}

func (lir linear) Mult(i int) Range {
	lir.Max *= i
	lir.Min *= i
	return lir
}