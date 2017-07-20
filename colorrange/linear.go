package colorrange

import (
	"image/color"

	"github.com/200sc/go-dist/intrange"
)

// Linear color ranges return colors on a linear distribution
type Linear struct {
	r, g, b, a intrange.Range
}

// NewLinear returns a linear color distribution between min and maxColor
func NewLinear(minColor, maxColor color.Color) Range {
	r, g, b, a := minColor.RGBA()
	r2, g2, b2, a2 := maxColor.RGBA()
	return Linear{
		intrange.NewLinear(int(r), int(r2)),
		intrange.NewLinear(int(g), int(g2)),
		intrange.NewLinear(int(b), int(b2)),
		intrange.NewLinear(int(a), int(a2)),
	}
}

// EnforceRange rounds the input color's components so that they fall in the
// given range.
func (l Linear) EnforceRange(c color.Color) color.Color {
	r3, g3, b3, a3 := c.RGBA()
	r4 := l.r.EnforceRange(int(r3))
	g4 := l.g.EnforceRange(int(g3))
	b4 := l.b.EnforceRange(int(b3))
	a4 := l.a.EnforceRange(int(a3))
	return color.RGBA64{uint16(r4), uint16(g4), uint16(b4), uint16(a4)}
}

// InRange returns whether the input color falls in this range
func (l Linear) InRange(c color.Color) bool {
	r3, g3, b3, a3 := c.RGBA()
	return l.r.InRange(int(r3)) &&
		l.g.InRange(int(g3)) &&
		l.b.InRange(int(b3)) &&
		l.a.InRange(int(a3))
}

// Poll returns a randomly chosen color in the bounds of this color range
func (l Linear) Poll() color.Color {
	r3 := l.r.Poll()
	g3 := l.g.Poll()
	b3 := l.b.Poll()
	a3 := l.a.Poll()
	return color.RGBA64{uint16(r3), uint16(g3), uint16(b3), uint16(a3)}
}
