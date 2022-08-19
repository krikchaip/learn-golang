package clockface

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	Precision        uint8   = 3
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
	<circle cx="%.[1]*[2]f" cy="%.[1]*[3]f" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

	<!-- hour hand -->
	<line x1="%.[1]*[2]f" y1="%.[1]*[3]f" x2="%.[1]*[8]f" y2="%.[1]*[9]f"
				style="fill:none;stroke:#000;stroke-width:7px;"/>

	<!-- minute hand -->
	<line x1="%.[1]*[2]f" y1="%.[1]*[3]f" x2="%.[1]*[6]f" y2="%.[1]*[7]f"
				style="fill:none;stroke:#000;stroke-width:7px;"/>

	<!-- second hand -->
	<line x1="%.[1]*[2]f" y1="%.[1]*[3]f" x2="%.[1]*[4]f" y2="%.[1]*[5]f"
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

func MakeHands(t time.Time) (
	hourHand Point,
	minuteHand Point,
	secondHand Point,
) {
	h, m, s := clockTimeInRadian(t)

	unitY, unitX := math.Sincos(s)
	secondHand = Point{unitX, unitY}.ShiftLength(SecondHandLength)

	unitY, unitX = math.Sincos(m)
	minuteHand = Point{unitX, unitY}.ShiftLength(MinuteHandLength)

	unitY, unitX = math.Sincos(h)
	hourHand = Point{unitX, unitY}.ShiftLength(HourHandLength)

	return
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
	h = float64(t.Hour()%12)/12 + m/12

	return
}

func ClockSVG(t time.Time) string {
	hourHand, minuteHand, secondHand := MakeHands(t)
	return fmt.Sprintf(
		strings.TrimLeft(tag, "\n "),
		Precision,
		OriginX, OriginY,
		secondHand.X, secondHand.Y,
		minuteHand.X, minuteHand.Y,
		hourHand.X, hourHand.Y,
	)
}
