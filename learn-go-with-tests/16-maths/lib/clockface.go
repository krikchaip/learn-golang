package clockface

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	OriginX          float64 = 150
	OriginY          float64 = 150
	SecondHandLength float64 = 90
	MinuteHandLength float64 = 80
	HourHandLength   float64 = 40
)

const tag = `
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
		width="100%%"
		height="100%%"
		viewBox="0 0 300 300"
		version="2.0">

	<!-- bezel -->
	<circle cx="%.3[1]f" cy="%.3[2]f" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

	<!-- hour hand -->
	<!-- <line x1="%.3[1]f" y1="%.3[2]f" x2="" y2=""
				style="fill:none;stroke:#000;stroke-width:7px;"/> -->

	<!-- minute hand -->
	<line x1="%.3[1]f" y1="%.3[2]f" x2="%.3[5]f" y2="%.3[6]f"
				style="fill:none;stroke:#000;stroke-width:7px;"/>

	<!-- second hand -->
	<line x1="%.3[1]f" y1="%.3[2]f" x2="%.3[3]f" y2="%.3[4]f"
				style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
`

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

// Shift a unit vector p by Origin constants and a specified Length l
func (p Point) ShiftLength(l float64) Point {
	p.X = OriginX + l*p.X
	p.Y = OriginY + l*p.Y
	return p
}

func (a Point) RoughlyEqual(b Point) bool {
	const threshold = 1e-7
	eq := func(a, b float64) bool {
		return math.Abs(a-b) < threshold
	}
	return eq(a.X, b.X) && eq(a.Y, b.Y)
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	_, _, radian := clockTimeInRadian(t)
	unitY, unitX := math.Sincos(radian)
	return Point{unitX, unitY}.ShiftLength(SecondHandLength)
}

func MinuteHand(t time.Time) Point {
	_, radian, _ := clockTimeInRadian(t)
	unitY, unitX := math.Sincos(radian)
	return Point{unitX, unitY}.ShiftLength(MinuteHandLength)
}

func clockTimeInRadian(t time.Time) (h, m, s float64) {
	radian := func(x float64) float64 {
		return 2*math.Pi*x - math.Pi/2
	}

	// convert to radian at the end of the function
	defer func() {
		s = radian(s)
		m = radian(m)
		h = radian(h)
	}()

	s = float64(t.Second()) / 60
	m = float64(t.Minute())/60 + s/60
	h = float64(t.Hour()%12)/12 + m/60

	return
}

func ClockSVG(t time.Time) string {
	secondHand := SecondHand(t)
	minuteHand := MinuteHand(t)
	return fmt.Sprintf(
		strings.TrimLeft(tag, "\n "),
		OriginX, OriginY,
		secondHand.X, secondHand.Y,
		minuteHand.X, minuteHand.Y,
	)
}
